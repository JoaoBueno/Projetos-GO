import java.io.Serializable;

public class KeyDescription implements Serializable {
  public final String rcsid = "$Id: KeyDescription.java 13950 2012-05-30 09:11:00Z marco_319 $";
  
  public static final int LEN_DUP = 1;
  
  public static final int LEN_SEG_NUM = 2;
  
  public static final int LEN_SEG_SIZE = 3;
  
  public static final int LEN_SEG_OFFS = 10;
  
  private byte[] stringDescr;
  
  public boolean duplicates;
  
  public Segs[] segment;
  
  public class Segs implements Serializable {
    public int size;
    
    public int offset;
    
    public boolean equals(Segs cmp) {
      return (this.size == cmp.size && this.offset == cmp.offset);
    }
    
    public byte[] toByteArray() {
      return IOUtil.int2Csv(new int[] { this.size, this.offset }, new int[] { 3, 10 });
    }
    
    public String toString() {
      return new String(toByteArray());
    }
  }
  
  public KeyDescription(byte[] descr) {
    int[] res = IOUtil.csv2Int(descr);
    int[] len = new int[res.length];
    len[0] = 2;
    len[1] = 1;
    for (int i = 2; i < len.length - 1; ) {
      len[i++] = 3;
      len[i++] = 10;
    } 
    this.stringDescr = IOUtil.int2Csv(res, len);
  }
  
  public KeyDescription(int nSeg, boolean dup) {
    this.duplicates = dup;
    this.segment = new Segs[nSeg];
    for (int i = 0; i < nSeg; i++)
      this.segment[i] = new Segs(); 
  }
  
  public boolean isDup() {
    int nSeg = getNumSegments();
    return this.duplicates;
  }
  
  public void setSegment(int n, int len, int offs) {
    int nSeg = getNumSegments();
    if (n < nSeg) {
      (this.segment[n]).size = len;
      (this.segment[n]).offset = offs;
    } 
  }
  
  public Segs getSegment(int n) {
    Segs Return;
    int nSeg = getNumSegments();
    if (n < nSeg) {
      Return = this.segment[n];
    } else {
      Return = null;
    } 
    return Return;
  }
  
  public int getNumSegments() {
    int Return;
    if (this.segment == null) {
      if (this.stringDescr != null) {
        int[] res = IOUtil.csv2Int(this.stringDescr);
        if (res.length >= 4) {
          this.duplicates = (res[1] == 1);
          this.segment = new Segs[(res.length - 2) / 2];
          for (int i = 0, j = 2; i < this.segment.length; i++) {
            this.segment[i] = new Segs();
            (this.segment[i]).size = res[j++];
            (this.segment[i]).offset = res[j++];
          } 
          Return = this.segment.length;
        } else {
          Return = 0;
        } 
      } else {
        Return = 0;
      } 
    } else {
      Return = this.segment.length;
    } 
    return Return;
  }
  
  public byte[] toByteArray() {
    if (this.stringDescr == null) {
      int[] values = new int[this.segment.length * 2 + 2];
      int[] lengths = new int[values.length];
      values[0] = this.segment.length;
      lengths[0] = 2;
      values[1] = this.duplicates ? 1 : 0;
      lengths[1] = 1;
      for (int i = 0, j = 2; i < this.segment.length; i++) {
        values[j] = (this.segment[i]).size;
        lengths[j++] = 3;
        values[j] = (this.segment[i]).offset;
        lengths[j++] = 10;
      } 
      this.stringDescr = IOUtil.int2Csv(values, lengths);
    } 
    return this.stringDescr;
  }
  
  public String toString() {
    return new String(toByteArray());
  }
  
  public int hashCode() {
    return toString().hashCode();
  }
  
  public int length() {
    int Return = 0;
    for (int i = getNumSegments() - 1; i >= 0; i--)
      Return += (this.segment[i]).size; 
    return Return;
  }
  
  public boolean equals(KeyDescription cmp) {
    int nSegs;
    if ((nSegs = getNumSegments()) == cmp.getNumSegments()) {
      for (int i = 0; i < nSegs; i++) {
        if (!this.segment[i].equals(cmp.segment[i]))
          return false; 
      } 
      return true;
    } 
    return false;
  }
  
  public static void main(String[] argv) {
    KeyDescription k1 = new KeyDescription(argv[0].getBytes());
    int nSeg = k1.getNumSegments();
    boolean dup = k1.isDup();
    System.out.println("nseg=" + nSeg + ", isDup=" + dup);
    Segs[] s1 = new Segs[nSeg];
    for (int i = 0; i < nSeg; i++) {
      s1[i] = k1.getSegment(i);
      System.out.println("seg=" + i + ", =[" + s1[i].toString() + "]");
    } 
    KeyDescription k2 = new KeyDescription(nSeg, dup);
    for (int j = 0; j < nSeg; j++)
      k2.setSegment(j, (s1[j]).size, (s1[j]).offset); 
    System.out.println("[" + k2.toString() + "]");
    System.out.println("->" + k2.equals(k1) + "<-");
  }
}
