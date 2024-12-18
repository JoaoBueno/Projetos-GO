IDENTIFICATION DIVISION.
PROGRAM-ID.    PARSETST.

ENVIRONMENT DIVISION.
INPUT-OUTPUT SECTION.
FILE-CONTROL.

    SELECT ARQTXT ASSIGN "arquivo.json"
           ORGANIZATION LINE SEQUENTIAL
           FILE STATUS SW-STATUS.

DATA DIVISION.
FILE SECTION.

FD  ARQTXT.
01  TXT-LINHA                   PIC X(65000).

WORKING-STORAGE SECTION.
77  SW-STATUS                   PIC X(002)          VALUE SPACES.
77  WS-PARA                     PIC X(001)          VALUE SPACES.
77  RET                         SIGNED-LONG.
77  WP-XFD-JSON                 POINTER.
77  WA-XFD-JSON                 PIC X(65000).
77  WA-MD5                      PIC X(032).
77                              PIC X(001)          VALUE X"00".
77  WA-AQRUIVO                  PIC X(032).
77                              PIC X(001)          VALUE X"00".
77  LEN                         SIGNED-LONG.

01  WS-DATA                     PIC 9(008)          VALUE  ZEROS.
01  REDEFINES WS-DATA.
    03  WS-DIA                  PIC 9(002).
    03  WS-MES                  PIC 9(002).
    03  WS-ANO                  PIC 9(004).

01  WS-DATA1                    PIC 9(008)          VALUE  ZEROS.
01  REDEFINES WS-DATA1.
    03  WS-DIA1                 PIC 9(002).
    03  WS-MES1                 PIC 9(002).
    03  WS-ANO1                 PIC 9(004).

01  WA-XFD-PARSE.
    03  XFD-NAME                PIC X(030).
    03  XFD-OFFSET              PIC 9(005).
    03  XFD-LENGTH              PIC 9(005).
    03  XFD-TYPE                PIC 9(005).
    03  XFD-DIGITS              PIC 9(005).
    03  XFD-SCALE               PIC S9(005).
    03  XFD-USER-TYPE           PIC 9(005).
    03  XFD-CONDITION           PIC 9(005).
    03  XFD-LEVEL               PIC 9(005).
    03  XFD-FORMAT              PIC X(030).

copy "bacu/acugui.def".
copy "bacu/parsexfd.ws".

PROCEDURE DIVISION.
PROCED-00.
    DISPLAY OMITTED BLANK SCREEN COLOR 1.
    SET CONFIGURATION "DLL-CONVENTION" TO 0.

    move "/multidad/xfd/aivenfcp.xfd" to xfdfile.
    move "/multidad/vendas/aivenfcp" to filename.

    CALL "./libmd5.so".
    CALL "./libparsexfd.so".

    call "parsexfd" using parse-xfd-op.
    if   parse-error
         display message box "Erro lendo a XFD" X"0A"
       	                     parsexfd-text-error-message(parse-flag),
       	         title "Erros na XFD" type is MB-OK, icon is MB-ERROR-ICON.

    initialize WA-XFD-JSON.
    move 0 to xfd-max-field-name-len.
    perform varying xfd-field-index from 1 by 1
			until xfd-field-index > xfd-total-number-fields
        call "parsexfd" using get-field-info-op
      *>  display message box "XFD-FIELD-NAME         " XFD-FIELD-NAME           X"0A"
      *>                      "XFD-FIELD-OFFSET       " XFD-FIELD-OFFSET         X"0A"
      *>                      "XFD-FIELD-LENGTH       " XFD-FIELD-LENGTH         X"0A"
      *>                      "XFD-FIELD-TYPE         " XFD-FIELD-TYPE           X"0A"
      *>                      "XFD-FIELD-DIGITS       " XFD-FIELD-DIGITS         X"0A"
      *>                      "XFD-FIELD-SCALE        " XFD-FIELD-SCALE          X"0A"
      *>                      "XFD-FIELD-USER-TYPE    " XFD-FIELD-USER-TYPE      X"0A"
      *>                      "XFD-FIELD-CONDITION    " XFD-FIELD-CONDITION      X"0A"
      *>                      "XFD-FIELD-LEVEL        " XFD-FIELD-LEVEL          X"0A"
      *>                      "XFD-FIELD-FORMAT       " XFD-FIELD-FORMAT         X"0A"
      *>                      "XFD-FIELD-OCCURS-DEPTH " XFD-FIELD-OCCURS-DEPTH   X"0A"
      
        initialize WA-XFD-PARSE
        MOVE XFD-FIELD-NAME      TO XFD-NAME
        MOVE XFD-FIELD-OFFSET    TO XFD-OFFSET
        MOVE XFD-FIELD-LENGTH    TO XFD-LENGTH
        MOVE XFD-FIELD-TYPE      TO XFD-TYPE
        MOVE XFD-FIELD-DIGITS    TO XFD-DIGITS
        MOVE XFD-FIELD-SCALE     TO XFD-SCALE
        MOVE XFD-FIELD-USER-TYPE TO XFD-USER-TYPE
        MOVE XFD-FIELD-CONDITION TO XFD-CONDITION
        MOVE XFD-FIELD-LEVEL     TO XFD-LEVEL
        MOVE XFD-FIELD-FORMAT    TO XFD-FORMAT

        string WA-XFD-JSON   DELIMITED BY "   "
               XFD-NAME      DELIMITED BY "   "
               ","           DELIMITED BY SIZE
               XFD-OFFSET    DELIMITED BY SIZE
               ","           DELIMITED BY SIZE
               XFD-LENGTH    DELIMITED BY SIZE
               ","           DELIMITED BY SIZE
               XFD-TYPE      DELIMITED BY SIZE
               ","           DELIMITED BY SIZE
               XFD-DIGITS    DELIMITED BY SIZE
               ","           DELIMITED BY SIZE
               XFD-SCALE     DELIMITED BY SIZE
               ","           DELIMITED BY SIZE
               XFD-USER-TYPE DELIMITED BY SIZE
               ","           DELIMITED BY SIZE
               XFD-CONDITION DELIMITED BY SIZE
               ","           DELIMITED BY SIZE
               XFD-LEVEL     DELIMITED BY SIZE
               ","           DELIMITED BY SIZE
               XFD-FORMAT    DELIMITED BY "   "
               X"0A"         DELIMITED BY SIZE
               INTO WA-XFD-JSON

*        CALL "XFDParse" USING BY REFERENCE WA-XFD-PARSE
*                              BY REFERENCE RET
*                        GIVING RETURN-CODE
*        END-CALL
    end-perform.

    string WA-XFD-JSON  DELIMITED BY "   "
           X"00"        DELIMITED BY SIZE
           INTO WA-XFD-JSON.

    CALL "MD5String"
             USING BY REFERENCE WA-XFD-JSON
                   BY REFERENCE WP-XFD-JSON
                   BY REFERENCE LEN
             GIVING RETURN-CODE
    END-CALL.

    CALL "C$MEMCPY" USING BY REFERENCE WA-MD5, BY VALUE WP-XFD-JSON, LEN.

    STOP " ".

    MOVE "aivencfp" to WA-AQRUIVO.

    CALL "XFDCheck"
             USING BY REFERENCE WA-AQRUIVO
                   BY REFERENCE WA-MD5
                   BY REFERENCE RET | 0 - OK 
             GIVING RETURN-CODE
    END-CALL.




    OPEN OUTPUT ARQTXT.
    WRITE TXT-LINHA FROM WA-XFD-JSON.
    CLOSE ARQTXT.


    CALL "XFDP" USING BY REFERENCE WA-XFD-JSON
                      BY REFERENCE RET
                GIVING RETURN-CODE
    END-CALL.



    STOP RUN.



    CALL "XFDtoJson" USING BY REFERENCE WP-XFD-JSON
                           BY REFERENCE LEN
    END-CALL.

    CALL "C$MEMCPY" USING BY REFERENCE WA-XFD-JSON, BY VALUE WP-XFD-JSON, LEN.

    OPEN OUTPUT ARQTXT.
    WRITE TXT-LINHA FROM WA-XFD-JSON.
    CLOSE ARQTXT.

    CALL "XFDCreateTable".

    *> DISPLAY message box WA-XFD-JSON.
    accept ws-para.
