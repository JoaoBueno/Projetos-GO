import java.io.File;
import java.io.FileNotFoundException;
import java.io.IOException;

public class ScanVision implements DynamicFile, RuntimeErrorsNumbers, Cloneable {
  public final String rcsid = "$Id";
  
  private VisionFile theFile;
  
  private KeyDescription[] keys;
  
  private String path = "";
  
  private int errno;
  
  private Exception errio = null;
  
  public String getSysErrno() {
    if (this.errio != null)
      return this.errio.getMessage(); 
    return "";
  }
  
  public String getErrMsg() {
    if (this.errio != null)
      return this.errio.getMessage(); 
    return "";
  }
  
  public int getCobErrno() {
    return this.errno;
  }
  
  private int mapError(int e) {
    this.errno = e;
    return 0;
  }
  
  public synchronized long getNumRecords() {
    if (this.theFile != null)
      return this.theFile.getValidRecordsNum(); 
    return -1L;
  }
  
  public synchronized int getNumKeys() {
    int Return;
    if (this.keys != null) {
      Return = this.keys.length;
    } else if (this.theFile != null) {
      this.keys = new KeyDescription[Return = this.theFile.nKeys];
    } else {
      Return = -1;
    } 
    return Return;
  }
  
  public int getMaxRecordSize() {
    if (this.theFile != null)
      return this.theFile.maxRec; 
    return -1;
  }
  
  public int getMinRecordSize() {
    if (this.theFile != null)
      return this.theFile.minRec; 
    return -1;
  }
  
  public synchronized KeyDescription getKey(int num) {
    KeyDescription Return;
    if (this.theFile != null && num >= 0 && num < getNumKeys()) {
      if (this.keys[num] == null) {
        this.keys[num] = new KeyDescription((this.theFile.keys[num]).nparts, (this.theFile.keys[num]).duplicates);
        for (int i = 0; i < (this.theFile.keys[num]).nparts; i++)
          this.keys[num].setSegment(i, ((this.theFile.keys[num]).parts[i]).length, ((this.theFile.keys[num]).parts[i]).offset); 
      } 
      Return = this.keys[num];
    } else {
      Return = null;
      mapError(102);
    } 
    return Return;
  }
  
  public byte[] getSequence() {
    return this.theFile.collatingSequence;
  }
  
  public int build(String ath, String comment, int blockingFactor, int preAllocate, int extensionFactor, int compressionFactor, int ecryptionFlag, int maxRecordSize, int minRecordSize, KeyDescription[] keys, byte[] collating, boolean assignExt) {
    return mapError(126);
  }
  
  public synchronized int open(String pth, int openType, int lockType, KeyDescription[] k, int maxRec, int minRec, int nKeys, int accessMode, boolean optional, boolean assignExt) {
    this.path = pth;
    // if (assignExt)
    //   this.path = BaseFile.expandFileName(this.path); 
    try {
      this.theFile = new VisionFile(this.path);
    } catch (FileNotFoundException ex) {
      if ((new File(this.path)).exists())
        return mapError(131); 
      return mapError(130);
    } catch (IOException ex) {
      this.errio = ex;
      return mapError(133);
    } 
    return 1;
  }
  
  public synchronized void setCurrentRecord(long nRec) {}
  
  public synchronized long getCurrentRecord() {
    return mapError(126);
  }
  
  public String getDescription() {
    return this.path;
  }
  
  public synchronized int close() {
    if (this.theFile != null) {
      this.theFile.close();
      this.theFile = null;
      this.keys = null;
      this.path = "";
    } 
    return 1;
  }
  
  public boolean isOpen() {
    return (this.theFile != null);
  }
  
  public synchronized long write(byte[] record, int offs, int size, boolean lock) {
    return mapError(126);
  }
  
  public synchronized long rewrite(byte[] record, int offs, int size, boolean lock) {
    return mapError(126);
  }
  
  public synchronized long delete(byte[] record, int offs) {
    return mapError(126);
  }
  
  public synchronized long next(byte[] record, int offs, int lock) {
    int Return;
    if (this.theFile != null) {
      Return = this.theFile.readNext(record, offs, record.length - offs);
      if (Return < 0)
        Return = mapError(110); 
    } else {
      Return = mapError(101);
    } 
    return Return;
  }
  
  public synchronized long previous(byte[] record, int offs, int lock) {
    return mapError(126);
  }
  
  public long read(byte[] record, int offs, KeyDescription k, int lock) {
    return mapError(126);
  }
  
  public long read(byte[] record, int offs, int keyNum, int lock) {
    return mapError(126);
  }
  
  public long start(byte[] record, int offs, int kNum, int kSize, int mode) {
    return mapError(126);
  }
  
  public long start(byte[] record, int offs, KeyDescription k, int kSize, int mode) {
    return mapError(126);
  }
  
  public synchronized int unlock() {
    return mapError(126);
  }
  
  public synchronized int recover() {
    return mapError(126);
  }
  
  public void sync(int allFiles) {}
  
  public int remove(String name) {
    return mapError(126);
  }
  
  public int rename(String src, String dst) {
    return mapError(126);
  }
  
  public int begin() {
    return 1;
  }
  
  public int commit(int ctx) {
    return 1;
  }
  
  public int rollback() {
    return 1;
  }
  
  public String getVersion() {
    return "VisionScan v" + this.theFile.vVersion;
  }
  
  public boolean isKeySelectedByNum() {
    return true;
  }
  
  // public void finalize() {
  //   close();
  // }
  
  public static void main(String[] argv) throws Exception {
    VisionFile vf = new VisionFile(argv[0]);
    vf.printInfo();
  }
}
