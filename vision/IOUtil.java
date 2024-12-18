public class IOUtil {
  public final String rcsid = "$Id: IOUtil.java 13950 2012-05-30 09:11:00Z marco_319 $";
  
  public static byte[] int2Csv(int[] in, int[] len) {
    int totLen = in.length - 1;
    int i;
    for (i = 0; i < in.length; i++)
      totLen += len[i]; 
    byte[] Return = new byte[totLen];
    i = 0;
    int j = 0;
    while (true) {
      int tmpNum = in[i];
      for (int k = len[i] - 1; k >= 0; k--) {
        Return[j--] = (byte)(48 + tmpNum % 10);
        tmpNum /= 10;
      } 
      j += len[i] + 1;
      if (++i < in.length) {
        Return[j++] = 44;
        continue;
      } 
      break;
    } 
    return Return;
  }
  
  public static int[] csv2Int(byte[] in) {
    int i, j;
    for (i = 0, j = 1; i < in.length; i++) {
      switch (in[i]) {
        case 44:
          j++;
          break;
      } 
    } 
    int[] Return = new int[j];
    for (i = 0, j = 0; i < in.length; i++) {
      switch (in[i]) {
        case 44:
          j++;
          break;
        case 48:
          Return[j] = Return[j] * 10;
          break;
        case 49:
          Return[j] = Return[j] * 10;
          Return[j] = Return[j] + 1;
          break;
        case 50:
          Return[j] = Return[j] * 10;
          Return[j] = Return[j] + 2;
          break;
        case 51:
          Return[j] = Return[j] * 10;
          Return[j] = Return[j] + 3;
          break;
        case 52:
          Return[j] = Return[j] * 10;
          Return[j] = Return[j] + 4;
          break;
        case 53:
          Return[j] = Return[j] * 10;
          Return[j] = Return[j] + 5;
          break;
        case 54:
          Return[j] = Return[j] * 10;
          Return[j] = Return[j] + 6;
          break;
        case 55:
          Return[j] = Return[j] * 10;
          Return[j] = Return[j] + 7;
          break;
        case 56:
          Return[j] = Return[j] * 10;
          Return[j] = Return[j] + 8;
          break;
        case 57:
          Return[j] = Return[j] * 10;
          Return[j] = Return[j] + 9;
          break;
        default:
          break;
      } 
    } 
    return Return;
  }
  
  public static final boolean equals(byte[] src, int srcPos, byte[] dst, int dstPos, int len) throws ArrayIndexOutOfBoundsException {
    for (int end = srcPos + len; srcPos < end; srcPos++, dstPos++) {
      if (dst[dstPos] != src[srcPos])
        return false; 
    } 
    return true;
  }
}
