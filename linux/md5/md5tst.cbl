IDENTIFICATION DIVISION.
PROGRAM-ID.      MD5TST.
*------------------------------------------------------------------------------*
* GERA UM MD5                                                                  *
* - TODOS                                                                      *
*                                                                              *
* CRIACAO...: 02/01/2017 - V6.00.000 - FBS                                     *
* ALTERACAO.:                                                                  *
*                                                                              *
* CODIGO FONTE DA BIBLIOTECA /fontes/delphi/dll/md5/md5.c                      *
* TEM QUE CRIAR A DLL NO LINUX E NO WINDOWS.                                   *
*                                                                              *
*                                                                              *
*------------------------------------------------------------------------------*
ENVIRONMENT DIVISION.
CONFIGURATION SECTION.
SPECIAL-NAMES.
    DECIMAL-POINT IS COMMA.
INPUT-OUTPUT SECTION.
FILE-CONTROL.

DATA DIVISION.
FILE SECTION.

WORKING-STORAGE SECTION.

77  WAP-USAMD5                  PIC X(001).

77  LSISOP                      PIC X(001).

77  WA-LIB                      PIC X(150)              VALUE SPACES.

77  STR                         PIC X(100)              VALUE "m.jpg".
77                              PIC X(001)              VALUE X"00".
77  RET                         pointer.
77  RETO                        PIC X(100).
77                              PIC X(001)              VALUE X"00".
77  LEN                         SIGNED-LONG.
77  X                           SIGNED-LONG.
77  Y                           SIGNED-LONG.
77                              PIC X(001)              VALUE X"00".

PROCEDURE DIVISION.
RT00-00-INICIO.
    ACCEPT STR UPDATE.

* 0 - FUNCOES EM C
    SET CONFIGURATION "DLL-CONVENTION" TO 0.

    MOVE "./libmd5.so" TO WA-LIB | LINUX

    CALL WA-LIB.

    CALL "MD5String"
             USING BY REFERENCE STR
                   BY REFERENCE RET
                   BY REFERENCE X
             GIVING RETURN-CODE
    END-CALL.

    CALL "C$MEMCPY" USING BY REFERENCE reto, BY VALUE ret, x
    display str.
    display ret.
    display x.
    display retO.

    CALL "MD5File"
             USING BY REFERENCE STR
                   BY REFERENCE RET
                   BY REFERENCE X
             GIVING RETURN-CODE
    END-CALL.

    CALL "C$MEMCPY" USING BY REFERENCE reto, BY VALUE ret, x
    display str.
    display ret.
    display x.
    display retO.
    CANCEL WA-LIB.

    display "tecle enter para sair."
    accept y.

RT00-10-FINAL.
    EXIT PROGRAM.
    STOP RUN.
