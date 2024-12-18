public interface DynamicFile extends IOConstants {
  public static final String rcsid = "$Id: DynamicFile.java 18567 2014-08-26 15:37:07Z marco_319 $";
  
  int open(String paramString, int paramInt1, int paramInt2, KeyDescription[] paramArrayOfKeyDescription, int paramInt3, int paramInt4, int paramInt5, int paramInt6, boolean paramBoolean1, boolean paramBoolean2);
  
  boolean isOpen();
  
  String getDescription();
  
  int getCobErrno();
  
  String getSysErrno();
  
  String getErrMsg();
  
  int close();
  
  int build(String paramString1, String paramString2, int paramInt1, int paramInt2, int paramInt3, int paramInt4, int paramInt5, int paramInt6, int paramInt7, KeyDescription[] paramArrayOfKeyDescription, byte[] paramArrayOfbyte, boolean paramBoolean);
  
  int getNumKeys();
  
  int getMaxRecordSize();
  
  int getMinRecordSize();
  
  KeyDescription getKey(int paramInt);
  
  long getNumRecords();
  
  byte[] getSequence();
  
  void setCurrentRecord(long paramLong);
  
  long getCurrentRecord();
  
  long read(byte[] paramArrayOfbyte, int paramInt1, int paramInt2, int paramInt3);
  
  long read(byte[] paramArrayOfbyte, int paramInt1, KeyDescription paramKeyDescription, int paramInt2);
  
  long next(byte[] paramArrayOfbyte, int paramInt1, int paramInt2);
  
  long previous(byte[] paramArrayOfbyte, int paramInt1, int paramInt2);
  
  long start(byte[] paramArrayOfbyte, int paramInt1, int paramInt2, int paramInt3, int paramInt4);
  
  long start(byte[] paramArrayOfbyte, int paramInt1, KeyDescription paramKeyDescription, int paramInt2, int paramInt3);
  
  long write(byte[] paramArrayOfbyte, int paramInt1, int paramInt2, boolean paramBoolean);
  
  long rewrite(byte[] paramArrayOfbyte, int paramInt1, int paramInt2, boolean paramBoolean);
  
  long delete(byte[] paramArrayOfbyte, int paramInt);
  
  int unlock();
  
  int remove(String paramString);
  
  int rename(String paramString1, String paramString2);
  
  void sync(int paramInt);
  
  int begin();
  
  int commit(int paramInt);
  
  int rollback();
  
  int recover();
  
  String getVersion();
  
  boolean isKeySelectedByNum();
}
