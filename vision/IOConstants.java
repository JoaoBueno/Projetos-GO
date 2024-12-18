public interface IOConstants {
  public static final String rcsid = "$Id: IOConstants.java 21568 2016-03-18 14:37:37Z marco_319 $ (C) picoSoft";
  
  public static final int SUCCESS = 1;
  
  public static final int FAILURE = 0;
  
  public static final int WARNING = -1;
  
  public static final int ACCESS_SEQUENTIAL = 1;
  
  public static final int ACCESS_RANDOM = 2;
  
  public static final int ACCESS_DYNAMIC = 3;
  
  public static final int ACCESS_MASK = 15;
  
  public static final int ORG_INDEXED = 16;
  
  public static final int ORG_RELATIVE = 32;
  
  public static final int ORG_SEQUENTIAL = 48;
  
  public static final int ORG_LINE_SEQUENTIAL = 64;
  
  public static final int ORG_MASK = 240;
  
  public static final int OPEN_CLOSED = 0;
  
  public static final int OPEN_INPUT = 1;
  
  public static final int OPEN_OUTPUT = 2;
  
  public static final int OPEN_IO = 3;
  
  public static final int OPEN_EXTEND = 6;
  
  public static final int TRANSACTION = 8;
  
  public static final String[] OPEN = new String[] { "CLOSED", "INPUT", "OUTPUT", "IO", "", "", "EXTEND", "", "TRANSACTION" };
  
  public static final int LOCK_NONE = 0;
  
  public static final int LOCK_EXCLUSIVE = 1;
  
  public static final int LOCK_NO_OTHERS = 2;
  
  public static final int LOCK_READERS = 3;
  
  public static final int LOCK_WRITERS = 4;
  
  public static final int LOCK_UPDATERS = 5;
  
  public static final int LOCK_ALL = 6;
  
  public static final int LOCK_TRANS = 7;
  
  public static final String[] LOCK = new String[] { "NONE", "EXCLUSIVE", "NO_OTHERS", "READERS", "WRITERS", "UPDATERS", "ALL", "TRANS" };
  
  public static final int LOCK_AUTOMATIC = 128;
  
  public static final int LOCK_MULTI = 256;
  
  public static final int LOCK_MASS_UPDATE = 512;
  
  public static final int LOCK_BULK = 1024;
  
  public static final int LOCK_TRANSACTION = 2048;
  
  public static final int FILE_ENCRYPTED = 4096;
  
  public static final int FILE_KEY_COMP = 8192;
  
  public static final int FILE_DATA_COMP = 16384;
  
  public static final int ADVANCING_NONE = 0;
  
  public static final int ADVANCING_AFTER = 1;
  
  public static final int ADVANCING_BEFORE = 2;
  
  public static final int ADVANCING_AFTER_PAGE = 3;
  
  public static final int ADVANCING_BEFORE_PAGE = 4;
  
  public static final int ADVANCING_AFTER_CHNL = 5;
  
  public static final int ADVANCING_BEFORE_CHNL = 6;
  
  public static final int ADVANCING_ONLY = 7;
  
  public static final int ADVANCING_PAGE_ONLY = 8;
  
  public static final int START_FIRST = 0;
  
  public static final int START_LAST = 1;
  
  public static final int START_NEXT = 2;
  
  public static final int START_PREV = 3;
  
  public static final int START_CURR = 4;
  
  public static final int START_EQUAL = 5;
  
  public static final int START_GREAT = 6;
  
  public static final int START_GTEQ = 7;
  
  public static final int START_LESS = 8;
  
  public static final int START_LTEQ = 9;
  
  public static final String[] START = new String[] { "FIRST", "LAST", "NEXT", "PREV", "CURR", "EQUAL", "GREAT", "GTEQ", "LESS", "LTEQ" };
  
  public static final int READ_IGNORE = -1;
  
  public static final int READ_NO_LOCK = 0;
  
  public static final int READ_LOCK = 1;
  
  public static final int READ_KEPT = 2;
  
  public static final int READ_WAIT = 3;
  
  public static final int READ_LOCK_MASK = 15;
  
  public static final int OPTS_NONE = 0;
  
  public static final int OPTS_NO_CR = 1;
  
  public static final int OPTS_NO_STRIP = 2;
  
  public static final int INFO_LOGICAL_PARAMS = -1;
  
  public static final int INFO_PHYSICAL_PARAMS = -2;
  
  public static final int INFO_COMMENT = -3;
  
  public static final int INFO_RECORD_COUNT = -4;
  
  public static final int INFO_COLLATING_SEQUENCE = -5;
  
  public static final int INFO_LOCK_COUNT = -6;
  
  public static final int INFO_SEGMENT_COUNT = -7;
  
  public static final int INFO_SEGMENT_INFO = -8;
  
  public static final int INFO_FILE_SIZE = -9;
  
  public static final int INFO_VERSION_NUMBER = -10;
  
  public static final int FILE_NORMAL = 0;
  
  public static final int FILE_PIPE = 1;
  
  public static final int FILE_STDIN = 2;
  
  public static final int FILE_STDOUT = 3;
  
  public static final int FILE_STDERR = 4;
  
  public static final int FILE_MASK = 15;
  
  public static final int FILE_TEXT = 16;
  
  public static final int FILE_BINARY = 32;
  
  public static final int CLOSE_NORMAL = 0;
  
  public static final int CLOSE_NO_REWIND = 1;
}
