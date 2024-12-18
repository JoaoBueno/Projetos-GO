IDENTIFICATION  DIVISION.
PROGRAM-ID.    VERIF-MD5.
*------------------------------------------------------------------------------*
* VERIFICA OS MD5 DE UMA PASTA INDICADA                                        *
*                                                                              *
* CRIACAO...: 05/02/2020 - BUENO - V6.00.000                                   *
* ALTERACAO.:   /  /     -                                                     *
*                                                                              *
*------------------------------------------------------------------------------*
ENVIRONMENT DIVISION.
CONFIGURATION SECTION.
SPECIAL-NAMES.
    DECIMAL-POINT IS COMMA.
INPUT-OUTPUT SECTION.
FILE-CONTROL.
    SELECT LISTATXT ASSIGN "md5-lista.cobol"
           ORGANIZATION LINE SEQUENTIAL
           FILE STATUS SW-STATUS.

DATA DIVISION.
FILE SECTION.
FD  LISTATXT.
01  TXT-LINHA.
    03  TXT-ARQUIVO             PIC X(060).
    03  TXT-MD5                 PIC X(032).

WORKING-STORAGE SECTION.
77  SW-STATUS                   PIC X(002).
77  PATTERN                     PIC X(005)          VALUE "*.svg".
77  MYDIR                       USAGE HANDLE.
77  RET                         POINTER.
77  RETO                        PIC X(100).
77  LEN                         SIGNED-LONG.
77  X                           SIGNED-LONG.
77  WA-ARQUIVO                  PIC X(200).
77                              PIC X(001)              VALUE X"00".
77  DIRECTORY                   PIC X(050)          VALUE "/usr/share/icons/Adwaita/scalable/apps/".
77  FILENAME                    PIC X(128).

copy "acucobol.def".

PROCEDURE DIVISION.
PROCED-00.
    OPEN OUTPUT LISTATXT.
    IF   SW-STATUS(1:1) NOT = ZEROS
         DISPLAY MESSAGE BOX "ERRO AO ABRIR O ARQUIVO, STATUS " SW-STATUS
         STOP RUN.

MAIN.
    CALL "C$LIST-DIRECTORY" USING LISTDIR-OPEN, DIRECTORY, PATTERN.
    MOVE RETURN-CODE TO MYDIR.
    IF MYDIR = 0
       STOP RUN
    END-IF.

    SET CONFIGURATION "DLL-CONVENTION" TO 0.
    CALL "./libmd5.so".

    PERFORM WITH TEST AFTER UNTIL FILENAME = SPACES
            CALL "C$LIST-DIRECTORY" USING LISTDIR-NEXT, MYDIR, FILENAME
            INITIALIZE TXT-LINHA
            MOVE FILENAME TO TXT-ARQUIVO
            INITIALIZE WA-ARQUIVO
            STRING DIRECTORY DELIMITED BY " "
                   FILENAME  DELIMITED BY " "
                   INTO WA-ARQUIVO
            CALL "MD5File" USING BY REFERENCE WA-ARQUIVO
                                 BY REFERENCE RET
                                 BY REFERENCE X
                           GIVING RETURN-CODE
            END-CALL
            IF   X >= ZEROS
                 CALL "C$MEMCPY" USING BY REFERENCE reto, BY VALUE ret, x
            END-IF
            MOVE RETO TO TXT-MD5
            WRITE TXT-LINHA
    END-PERFORM.

    CALL "C$LIST-DIRECTORY" USING LISTDIR-CLOSE, MYDIR.
    STOP RUN.
