import java.io.EOFException;
import java.io.File;
import java.io.IOException;
import java.io.RandomAccessFile;

class VisionFile {
  static class KeyPart {
    final int offset;
    
    final short length;
    
    KeyPart(int o, short l) {
      this.offset = o;
      this.length = l;
    }
  }
  
  static class Key {
    final int nparts;
    
    final boolean duplicates;
    
    final int length;
    
    final short boh;
    
    final VisionFile.KeyPart[] parts;
    
    Key(int n, boolean dups, int len, short b) {
      this.nparts = n;
      this.parts = new VisionFile.KeyPart[this.nparts];
      this.duplicates = dups;
      this.length = len;
      this.boh = b;
    }
  }
  
  private final String idxExt = ".vix";
  
  private File vFile;
  
  private RandomAccessFile vRAFile;
  
  final char vVersion;
  
  private final int blockingFactor;
  
  private final int blockSize;
  
  private final int blockSize_4;
  
  final long numOfRecords;
  
  private final long deletedRecords;
  
  private final long totOpens;
  
  private final long userCount;
  
  private final long segmentSize;
  
  final int maxRec;
  
  final int minRec;
  
  final int nKeys;
  
  final Key[] keys;
  
  private final boolean compressed;
  
  private boolean hasDuplicates;
  
  private byte[] buffer;
  
  private final byte[] node;
  
  private final byte[] dummyInt = new byte[4];
  
  private int nodePnt;
  
  private long validRecordsNum;
  
  private long deletedRecordsNum;
  
  private final String fName;
  
  private final int nDataSegments;
  
  private int currDataSegment;
  
  private final long firstDataBlock;
  
  final byte[] collatingSequence;
  
  private String blanks = "                                            ";
  
  private static final long readInt48(RandomAccessFile raf) throws IOException {
    long a = raf.readChar();
    long Return = raf.readInt() & 0xFFFFFFFFL;
    Return |= a << 32L;
    return Return;
  }
  
  public VisionFile(String fileName) throws IOException {
    this.blanks = "                                            ";
    int lastDot = fileName.lastIndexOf('.');
    if (lastDot > 0 && !(new File(fileName + ".vix")).exists()) {
      this.fName = fileName.substring(0, lastDot);
    } else {
      this.fName = fileName;
    } 
    this.vFile = new File(fileName);
    this.vRAFile = new RandomAccessFile(this.vFile, "r");
    try {
      int dataBlock, offset, keySize;
      long csAddr;
      int i, v = this.vRAFile.readInt();
      if (v != 269620246 && v != 269620249)
        throw new IOException("Unrecognized Vision file!"); 
      this.vVersion = this.vRAFile.readChar();
      this.blockingFactor = this.vRAFile.readChar();
      this.blockSize = this.blockingFactor * 512;
      this.blockSize_4 = this.blockSize - 4;
      switch (this.vVersion) {
        case '\003':
          this.node = new byte[this.blockSize];
          this.nDataSegments = 0;
          this.vRAFile.seek(24L);
          this.firstDataBlock = this.vRAFile.readInt() & 0xFFFFFFFEL;
          this.vRAFile.seek(40L);
          this.numOfRecords = this.vRAFile.readInt() & 0xFFFFFFFFL;
          this.deletedRecords = this.vRAFile.readInt() & 0xFFFFFFFFL;
          this.totOpens = 0L;
          this.vRAFile.seek(56L);
          csAddr = this.vRAFile.readInt() & 0xFFFFFFFFL;
          if (csAddr != 0L) {
            this.collatingSequence = new byte[256];
            this.vRAFile.seek(csAddr + 16L);
            this.vRAFile.read(this.collatingSequence);
          } else {
            this.collatingSequence = null;
          } 
          this.vRAFile.seek(64L);
          this.userCount = this.vRAFile.readInt() & 0xFFFFFFFFL;
          this.segmentSize = 0L;
          this.vRAFile.seek(96L);
          this.maxRec = this.vRAFile.readChar();
          this.minRec = this.vRAFile.readChar();
          this.buffer = new byte[this.maxRec];
          this.nKeys = this.vRAFile.readByte() & 0xFF;
          this.keys = new Key[this.nKeys];
          dataBlock = 1;
          this.compressed = ((this.vRAFile.readByte() & 0xFF) != 0);
          if ((this.vRAFile.readByte() & 0xFF) != 0)
            throw new IOException("Unsupported feature: encryption"); 
          offset = 160;
          keySize = 32;
          for (i = 0; i < this.nKeys; i++) {
            int nextOffs = offset + keySize * i;
            int nextEnd = (nextOffs + keySize) % 512;
            if (nextEnd != 0 && nextEnd <= keySize) {
              offset = dataBlock * 512;
              dataBlock++;
              offset -= keySize * i;
            } 
            this.vRAFile.seek((offset + keySize * i));
            int boh1 = this.vRAFile.readByte() & 0xFF;
            int boh2 = this.vRAFile.readInt();
            int nparts = this.vRAFile.readByte() & 0xFF;
            boolean dups = (this.vRAFile.readByte() != 0);
            this.hasDuplicates |= dups;
            int len = this.vRAFile.readByte() & 0xFF;
            this.keys[i] = new Key(nparts, dups, len, (short)0);
            for (int j = 0; j < nparts; j++) {
              short boh = (short)(this.vRAFile.readByte() & 0xFF);
              short plen = (short)(this.vRAFile.readByte() & 0xFF);
              int offs = this.vRAFile.readChar();
              (this.keys[i]).parts[j] = new KeyPart(offs, plen);
            } 
          } 
          if (this.firstDataBlock > 0L) {
            this.vRAFile.seek(this.firstDataBlock);
          } else {
            this.vRAFile.seek(this.vRAFile.length());
          } 
          return;
        case '\004':
          this.node = null;
          this.vRAFile.seek(16L);
          this.nDataSegments = this.vRAFile.readShort();
          this.vRAFile.seek(30L);
          this.firstDataBlock = this.vRAFile.readInt() & 0xFFFFFFFFL;
          this.vRAFile.seek(52L);
          this.numOfRecords = this.vRAFile.readInt() & 0xFFFFFFFFL;
          this.deletedRecords = this.vRAFile.readInt() & 0xFFFFFFFFL;
          this.vRAFile.seek(64L);
          this.totOpens = this.vRAFile.readInt() & 0xFFFFFFFFL;
          this.vRAFile.seek(68L);
          csAddr = this.vRAFile.readInt() & 0xFFFFFFFFL;
          if (csAddr != 0L) {
            this.collatingSequence = new byte[256];
            this.vRAFile.seek(csAddr + 16L);
            this.vRAFile.read(this.collatingSequence);
          } else {
            this.collatingSequence = null;
          } 
          this.vRAFile.seek(76L);
          this.userCount = this.vRAFile.readInt() & 0xFFFFFFFFL;
          this.vRAFile.seek(102L);
          this.segmentSize = this.vRAFile.readInt() & 0xFFFFFFFFL;
          this.vRAFile.seek(116L);
          this.maxRec = this.vRAFile.readChar();
          this.minRec = this.vRAFile.readChar();
          this.buffer = new byte[this.maxRec];
          this.nKeys = this.vRAFile.readByte() & 0xFF;
          this.keys = new Key[this.nKeys];
          dataBlock = 1;
          this.compressed = ((this.vRAFile.readByte() & 0xFF) != 0);
          if ((this.vRAFile.readByte() & 0xFF) != 0)
            throw new IOException("Unsupported feature: encryption"); 
          offset = 160;
          keySize = 59;
          for (i = 0; i < this.nKeys; i++) {
            int nextOffs = offset + keySize * i;
            int nextEnd = (nextOffs + keySize) % 512;
            if (nextEnd != 0 && nextEnd <= keySize) {
              offset = dataBlock * 512;
              dataBlock++;
              offset -= keySize * i;
            } 
            this.vRAFile.seek((offset + keySize * i));
            int boh1 = this.vRAFile.readByte() & 0xFF;
            int boh2 = this.vRAFile.readChar();
            int boh3 = this.vRAFile.readInt();
            int nparts = this.vRAFile.readByte() & 0xFF;
            boolean dups = (this.vRAFile.readByte() != 0);
            this.hasDuplicates |= dups;
            int len = this.vRAFile.readByte() & 0xFF;
            short boh = (short)(this.vRAFile.readByte() & 0xFF);
            this.keys[i] = new Key(nparts, dups, len, boh);
            for (int j = 0; j < nparts; j++) {
              short plen = (short)(this.vRAFile.readByte() & 0xFF);
              int offs = this.vRAFile.readChar();
              (this.keys[i]).parts[j] = new KeyPart(offs, plen);
            } 
          } 
          if (this.firstDataBlock > 0L) {
            this.vRAFile.seek(this.firstDataBlock);
          } else {
            this.vRAFile.seek(this.vRAFile.length());
          } 
          return;
        case '\005':
          this.vRAFile.seek(20L);
          this.nDataSegments = this.vRAFile.readShort();
          this.node = null;
          this.vRAFile.seek(34L);
          this.firstDataBlock = this.vRAFile.readInt() & 0xFFFFFFFFL;
          this.vRAFile.seek(62L);
          this.numOfRecords = this.vRAFile.readInt() & 0xFFFFFFFFL;
          this.deletedRecords = this.vRAFile.readInt() & 0xFFFFFFFFL;
          this.vRAFile.seek(82L);
          csAddr = this.vRAFile.readInt() & 0xFFFFFFFFL;
          if (csAddr != 0L) {
            this.collatingSequence = new byte[256];
            this.vRAFile.seek(csAddr + 16L);
            this.vRAFile.read(this.collatingSequence);
          } else {
            this.collatingSequence = null;
          } 
          this.vRAFile.seek(92L);
          this.userCount = this.vRAFile.readInt() & 0xFFFFFFFFL;
          this.totOpens = this.vRAFile.readInt() & 0xFFFFFFFFL;
          this.vRAFile.seek(122L);
          this.segmentSize = this.vRAFile.readInt() & 0xFFFFFFFFL;
          this.vRAFile.seek(149L);
          this.maxRec = this.vRAFile.readInt();
          this.minRec = this.vRAFile.readInt();
          this.buffer = new byte[this.maxRec];
          this.nKeys = this.vRAFile.readByte() & 0xFF;
          this.keys = new Key[this.nKeys];
          dataBlock = 1;
          this.compressed = ((this.vRAFile.readByte() & 0xFF) != 0);
          if ((this.vRAFile.readByte() & 0xFF) != 0)
            throw new IOException("Unsupported feature: encryption"); 
          offset = 194;
          keySize = 106;
          for (i = 0; i < this.nKeys; i++) {
            int nextOffs = offset + keySize * i;
            int nextEnd = (nextOffs + keySize) % this.blockSize;
            if (nextEnd != 0 && nextEnd <= keySize) {
              offset = dataBlock * this.blockSize;
              dataBlock++;
              offset -= keySize * i;
            } 
            this.vRAFile.seek((offset + keySize * i));
            int boh1 = this.vRAFile.readByte() & 0xFF;
            int boh2 = this.vRAFile.readChar();
            int boh3 = this.vRAFile.readInt();
            int nparts = this.vRAFile.readByte() & 0xFF;
            boolean dups = (this.vRAFile.readByte() != 0);
            this.hasDuplicates |= dups;
            int len = this.vRAFile.readByte() & 0xFF;
            this.keys[i] = new Key(nparts, dups, len, (short)0);
            for (int j = 0; j < nparts; j++) {
              short boh = (short)(this.vRAFile.readByte() & 0xFF);
              short plen = (short)(this.vRAFile.readByte() & 0xFF);
              int offs = this.vRAFile.readInt();
              (this.keys[i]).parts[j] = new KeyPart(offs, plen);
            } 
          } 
          if (this.firstDataBlock > 0L) {
            this.vRAFile.seek(this.firstDataBlock);
          } else {
            this.vRAFile.seek(this.vRAFile.length());
          } 
          return;
        case '\006':
          this.vRAFile.seek(20L);
          this.nDataSegments = this.vRAFile.readShort();
          this.node = null;
          this.vRAFile.seek(34L);
          this.firstDataBlock = readInt48(this.vRAFile);
          this.vRAFile.seek(62L);
          this.numOfRecords = this.vRAFile.readInt() & 0xFFFFFFFFL;
          this.deletedRecords = this.vRAFile.readInt() & 0xFFFFFFFFL;
          this.vRAFile.seek(82L);
          csAddr = this.vRAFile.readInt() & 0xFFFFFFFFL;
          if (csAddr != 0L) {
            this.collatingSequence = new byte[256];
            this.vRAFile.seek(csAddr + 16L);
            this.vRAFile.read(this.collatingSequence);
          } else {
            this.collatingSequence = null;
          } 
          this.vRAFile.seek(92L);
          this.userCount = this.vRAFile.readInt() & 0xFFFFFFFFL;
          this.totOpens = readInt48(this.vRAFile);
          this.vRAFile.seek(122L);
          this.segmentSize = 0L;
          this.vRAFile.seek(134L);
          this.maxRec = this.vRAFile.readInt();
          this.minRec = this.vRAFile.readInt();
          this.buffer = new byte[this.maxRec];
          this.nKeys = this.vRAFile.readByte() & 0xFF;
          this.keys = new Key[this.nKeys];
          dataBlock = 1;
          this.compressed = ((this.vRAFile.readByte() & 0xFF) != 0);
          if ((this.vRAFile.readByte() & 0xFF) != 0)
            throw new IOException("Unsupported feature: encryption"); 
          offset = 178;
          keySize = 106;
          for (i = 0; i < this.nKeys; i++) {
            int nextOffs = offset + keySize * i;
            int nextEnd = (nextOffs + keySize) % this.blockSize;
            if (nextEnd != 0 && nextEnd <= keySize) {
              offset = dataBlock * this.blockSize;
              dataBlock++;
              offset -= keySize * i;
            } 
            this.vRAFile.seek((offset + keySize * i));
            int boh1 = this.vRAFile.readByte() & 0xFF;
            int boh2 = this.vRAFile.readInt();
            int boh3 = this.vRAFile.readChar();
            int nparts = this.vRAFile.readByte() & 0xFF;
            boolean dups = (this.vRAFile.readByte() != 0);
            this.hasDuplicates |= dups;
            int len = this.vRAFile.readByte() & 0xFF;
            this.keys[i] = new Key(nparts, dups, len, (short)0);
            for (int j = 0; j < nparts; j++) {
              short boh = (short)(this.vRAFile.readByte() & 0xFF);
              short plen = (short)(this.vRAFile.readByte() & 0xFF);
              int offs = this.vRAFile.readInt();
              (this.keys[i]).parts[j] = new KeyPart(offs, plen);
            } 
          } 
          if (this.firstDataBlock > 0L) {
            this.vRAFile.seek(this.firstDataBlock);
          } else {
            this.vRAFile.seek(this.vRAFile.length());
          } 
          return;
      } 
      throw new IOException("Vision version " + this.vVersion + " is not supported yet");
    } catch (IOException _ex) {
      close();
      throw _ex;
    } 
  }
  
  String fmt(String Return, int len, boolean left) {
    int sLen = Return.length();
    if (sLen < len)
      if (left) {
        Return = Return + "                                            ".substring(0, len - sLen);
      } else {
        Return = "                                            ".substring(0, len - sLen) + Return;
      }  
    return Return;
  }
  
  String fmt(String s, int len) {
    return fmt(s, len, false);
  }
  
  String fmt(long n, int len, boolean left) {
    String s = "" + n;
    return fmt(s, len, left);
  }
  
  String fmt(long n, int len) {
    return fmt(n, len, false);
  }
  
  public void printInfo() {
    System.out.println(this.vFile.getPath() + "  [vision version " + this.vVersion + "]");
    System.out.println("");
    File xFile = new File(this.vFile.getPath() + ".vix");
    System.out.println("# of records:" + fmt(this.numOfRecords, 18));
    System.err.println("# of deleted records:" + fmt(this.deletedRecords, 10));
    if (this.vVersion > '\003') {
      System.out.println("file size: " + fmt(this.vFile.length(), 20) + " (" + this.vFile.getPath() + ")");
      System.out.println("file size: " + fmt(xFile.length(), 20) + " (" + xFile.getPath() + ")");
      System.out.println("total file size:" + fmt(this.vFile.length() + xFile.length(), 15));
      System.out.println("segment size:" + fmt(this.segmentSize, 18));
    } else {
      System.out.println("file size: " + fmt(this.vFile.length(), 20));
    } 
    String c = this.compressed ? " compressed" : "";
    if (this.minRec == this.maxRec) {
      System.out.println("record size:" + fmt(this.maxRec, 19) + c);
    } else {
      System.out.println("record size (min/max):" + fmt(this.minRec, 9) + "/" + this.maxRec + c);
    } 
    System.out.println("# of keys: " + fmt(this.nKeys, 20));
    System.out.println("user count:" + fmt(this.userCount, 20));
    System.out.println("");
    System.out.println("Key  Dups    Seg-1     Seg-2     Seg-3     Seg-4     Seg-5     Seg-6");
    System.out.println("            (sz/of)   (sz/of)   (sz/of)   (sz/of)   (sz/of)   (sz/of)");
    System.out.println("");
    StringBuffer out = new StringBuffer();
    for (int i = 0; i < this.nKeys; i++) {
      out.delete(0, out.length());
      out.append(fmt(i, 3));
      out.append(fmt((this.keys[i]).duplicates ? "Y" : "N", 5));
      out.append("  ");
      for (int j = 0; j < (this.keys[i]).nparts; j++) {
        out.append(fmt(((this.keys[i]).parts[j]).length, 3));
        out.append("/");
        out.append(fmt(((this.keys[i]).parts[j]).offset, 6, true));
      } 
      System.out.println(out);
    } 
  }
  
  public long getValidRecordsNum() {
    return this.numOfRecords;
  }
  
  private int readNode3() throws IOException {
    int Return;
    if (this.nodePnt == 0 || this.nodePnt == this.blockSize_4) {
      this.nodePnt = 0;
      this.vRAFile.read(this.node, 0, this.blockSize_4);
      if (this.node[this.nodePnt++] == 2) {
        long nextNode = this.vRAFile.readInt() & 0xFFFFFFFFL;
        if (nextNode != 0L)
          this.vRAFile.seek(nextNode); 
        Return = this.blockSize_4 - 1;
      } else {
        throw new IOException("Corrupted file!!");
      } 
    } else {
      Return = this.blockSize_4 - this.nodePnt;
    } 
    return Return;
  }
  
  private int myReadChar3() throws IOException {
    int Return, rc = readNode3();
    if (rc < 1) {
      Return = -1;
    } else if (rc == 1) {
      Return = (this.node[this.nodePnt] & 0xFF) << 8;
      this.nodePnt = 0;
      if (readNode3() >= 1) {
        Return |= this.node[this.nodePnt++] & 0xFF;
      } else {
        Return = -1;
      } 
    } else {
      Return = (this.node[this.nodePnt++] & 0xFF) << 8;
      Return |= this.node[this.nodePnt++] & 0xFF;
    } 
    return Return;
  }
  
  private int myRead3(byte[] b, int offs, int len) throws IOException {
    int Return = len;
    int rc = readNode3();
    while (rc > 0 && rc < len) {
      System.arraycopy(this.node, this.nodePnt, b, offs, rc);
      offs += rc;
      len -= rc;
      this.nodePnt = 0;
      rc = readNode3();
    } 
    if (rc > 0) {
      System.arraycopy(this.node, this.nodePnt, b, offs, len);
      this.nodePnt += len;
    } else {
      Return = -1;
    } 
    return Return;
  }
  
  private int fillRecord(int iLen, boolean comp, byte[] out, int oOf, int oLen) {
    int Return;
    if (iLen == 0) {
      Return = 0;
    } else if (comp) {
      byte state = 0;
      int count = 0;
      byte chr = 0;
      Return = 0;
      for (int i = 0; i < iLen; i++) {
        switch (state) {
          case -5:
            count = this.buffer[i] & 0xFF;
            for (; count > 0; count--)
              out[Return++] = 0; 
            state = 0;
            break;
          case -4:
            count = this.buffer[i] & 0xFF;
            for (; count > 0; count--)
              out[Return++] = 48; 
            state = 0;
            break;
          case -3:
            count = this.buffer[i] & 0xFF;
            for (; count > 0; count--)
              out[Return++] = 32; 
            state = 0;
            break;
          case -2:
            if (chr == 0) {
              chr = this.buffer[i];
              break;
            } 
            count = this.buffer[i] & 0xFF;
            for (; count > 0; count--)
              out[Return++] = chr; 
            state = 0;
            chr = 0;
            break;
          case -6:
            out[Return++] = this.buffer[i];
            state = 0;
            break;
          default:
            switch (this.buffer[i]) {
              case -6:
              case -5:
              case -4:
              case -3:
              case -2:
                state = this.buffer[i];
                break;
            } 
            out[Return++] = this.buffer[i];
            break;
        } 
      } 
    } else {
      Return = Math.min(iLen, oLen);
      System.arraycopy(this.buffer, 0, out, oOf, Return);
    } 
    return Return;
  }
  
  private int readNext3(byte[] b, int offs, int len) {
    int Return = 0;
    while (Return == 0) {
      try {
        int recLen = myReadChar3();
        Return = myReadChar3();
        if (this.hasDuplicates)
          myRead3(this.dummyInt, 0, 4); 
        if (recLen > 0) {
          boolean comp;
          if (recLen > this.buffer.length)
            this.buffer = new byte[recLen]; 
          myRead3(this.buffer, 0, recLen);
          if (this.compressed) {
            if ((Return & 0x8000) != 0) {
              Return &= 0xFFFF7FFF;
              comp = false;
            } else {
              comp = true;
            } 
          } else {
            comp = false;
          } 
          if (Return == 0) {
            this.deletedRecordsNum++;
            continue;
          } 
          Return = fillRecord(Return, comp, b, offs, len);
          if (Return >= this.minRec && Return <= this.maxRec) {
            this.validRecordsNum++;
            continue;
          } 
          Return = -1;
        } 
      } catch (IOException _ex) {
        Return = -1;
      } 
    } 
    return Return;
  }
  
  private boolean newDataSegment() {
    if (this.currDataSegment < this.nDataSegments) {
      String ext;
      this.currDataSegment++;
      if (this.currDataSegment < 16) {
        ext = ".d0" + Integer.toHexString(this.currDataSegment);
      } else {
        ext = ".d" + Integer.toHexString(this.currDataSegment);
      } 
      try {
        this.vFile = new File(this.fName + ext);
        this.vRAFile.close();
        this.vRAFile = new RandomAccessFile(this.vFile, "r");
        this.vRAFile.seek(512L);
        return true;
      } catch (IOException _ex) {}
    } 
    return false;
  }
  
  private int readNext4(byte[] b, int offs, int len) {
    int Return = 0;
    while (Return == 0) {
      try {
        int recLen = this.vRAFile.readChar();
        Return = this.vRAFile.readChar();
        if (this.hasDuplicates)
          this.vRAFile.readInt(); 
        if (recLen > 0) {
          boolean comp;
          if (recLen > this.buffer.length)
            this.buffer = new byte[recLen]; 
          this.vRAFile.read(this.buffer, 0, recLen);
          if (this.compressed) {
            if ((Return & 0x8000) != 0) {
              Return &= 0xFFFF7FFF;
              comp = false;
            } else {
              comp = true;
            } 
          } else {
            comp = false;
          } 
          if (Return == 0) {
            this.deletedRecordsNum++;
            continue;
          } 
          Return = fillRecord(Return, comp, b, offs, len);
          if (Return >= this.minRec && Return <= this.maxRec) {
            this.validRecordsNum++;
            continue;
          } 
          Return = -1;
          continue;
        } 
        if (newDataSegment()) {
          Return = readNext4(b, offs, len);
          continue;
        } 
        Return = -1;
      } catch (EOFException _ex) {
        if (newDataSegment()) {
          Return = readNext4(b, offs, len);
          continue;
        } 
        Return = -1;
      } catch (IOException _ex) {
        Return = -1;
      } 
    } 
    return Return;
  }
  
  private int readNext5(byte[] b, int offs, int len) {
    int Return = 0;
    while (Return == 0) {
      try {
        int recLen = this.vRAFile.readInt();
        Return = this.vRAFile.readInt();
        byte deleted = this.vRAFile.readByte();
        int progr1 = this.vRAFile.readInt();
        char progr2 = this.vRAFile.readChar();
        if (recLen > 0) {
          boolean comp;
          if (recLen > this.buffer.length)
            this.buffer = new byte[recLen]; 
          this.vRAFile.read(this.buffer, 0, recLen);
          if (this.compressed && deleted != 4) {
            comp = true;
          } else {
            comp = false;
          } 
          if (deleted == 1) {
            this.deletedRecordsNum++;
            Return = 0;
            continue;
          } 
          Return = fillRecord(Return, comp, b, offs, len);
          if (Return >= this.minRec && Return <= this.maxRec) {
            this.validRecordsNum++;
            continue;
          } 
          Return = -1;
          continue;
        } 
        if (newDataSegment()) {
          Return = readNext5(b, offs, len);
          continue;
        } 
        Return = -1;
      } catch (EOFException _ex) {
        if (newDataSegment()) {
          Return = readNext5(b, offs, len);
          continue;
        } 
        Return = -1;
      } catch (IOException _ex) {
        Return = -1;
      } 
    } 
    return Return;
  }
  
  public int readNext(byte[] b, int offs, int len) {
    if (this.vVersion == '\003')
      return readNext3(b, offs, len); 
    if (this.vVersion == '\004')
      return readNext4(b, offs, len); 
    if (this.vVersion == '\005' || this.vVersion == '\006')
      return readNext5(b, offs, len); 
    return -1;
  }
  
  public int readNext(byte[] b) {
    return readNext(b, 0, b.length);
  }
  
  public void close() {
    if (this.vRAFile != null)
      try {
        this.vRAFile.close();
      } catch (IOException _ex) {} 
  }
}
