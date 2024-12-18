IDENTIFICATION DIVISION.
PROGRAM-ID.    t1_teste.

ENVIRONMENT DIVISION.

DATA DIVISION.
WORKING-STORAGE SECTION.
77  SW-STATUS                   PIC X(002)          VALUE SPACES.
77  WS-PARA                     PIC X(001)          VALUE SPACES.
77  RET                         SIGNED-LONG.
77  WA-XFD-JSON                 PIC X(100).
77  WP-XFD-JSON                 POINTER.
77  WA-MD5                      PIC X(032).
77  LEN                         SIGNED-LONG.

PROCEDURE DIVISION.
PROCED-00.
    DISPLAY OMITTED BLANK SCREEN COLOR 1.
    SET CONFIGURATION "DLL-CONVENTION" TO 0.

    CALL "./t1.so".

    move "teste" to WA-XFD-JSON.

    CALL "T1" USING BY REFERENCE WA-XFD-JSON
                    BY REFERENCE LEN
    END-CALL.

    move "teste" to WA-XFD-JSON.

    CALL "T2" USING BY REFERENCE WA-XFD-JSON
                    BY REFERENCE LEN
    END-CALL.

    move "teste" to WA-XFD-JSON.

    CALL "T3" USING BY REFERENCE WA-XFD-JSON
                    BY REFERENCE LEN
    END-CALL.

    CALL "T4" USING BY REFERENCE WA-XFD-JSON
                    BY REFERENCE LEN
    END-CALL.

    accept ws-para.
