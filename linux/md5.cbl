IDENTIFICATION DIVISION.
PROGRAM-ID.         MD5.
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

77  STR                         PIC X(100)              VALUE SPACES.
77                              PIC X(001)              VALUE X"00".
77  RETORNO                     PIC X(032)              VALUE SPACES.
77                              PIC X(001)              VALUE X"00".
77  LEN                         SIGNED-LONG.
77                              PIC X(001)              VALUE X"00".


    | COPY "/multifon/bacu/acucobol.def".

|LINKAGE SECTION.
|    COPY "/multifon/bok/md5.bok".
01  WAL-MD5.
    03  WAL-MD5-TEXTO            PIC X(100)             VALUE SPACES.
    03                           PIC X(001) VALUE X"00".
    03  WAL-MD5-RETORNO          PIC X(032)             VALUE SPACES.
    03                           PIC X(001) VALUE X"00".

SCREEN SECTION.

PROCEDURE DIVISION. 
RT00-00-INICIO.
    ACCEPT WAL-MD5-TEXTO.

    ACCEPT WAP-USAMD5     FROM ENVIRONMENT "USA-MD5".

    IF  FUNCTION UPPER-CASE(WAP-USAMD5) = "N"
        MOVE WAL-MD5 TO WAL-MD5-RETORNO
        GO TO RT00-10-FINAL
    END-IF.

    MOVE SPACES TO WAL-MD5-RETORNO.
    MOVE SPACES TO STR.
    MOVE SPACES TO RETORNO.
    MOVE ZEROS  TO LEN.

* 0 - FUNCOES EM C
    SET CONFIGURATION "DLL-CONVENTION" TO 0.

    MOVE "./md5lib.so"  TO WA-LIB | LINUX

    CALL WA-LIB.

    MOVE WAL-MD5-TEXTO TO STR.

    INSPECT STR TALLYING LEN FOR CHARACTERS BEFORE INITIAL "  ".

    CALL "retornamd5"
             USING BY REFERENCE STR, RETORNO
             BY VALUE LEN
             GIVING RETURN-CODE
    END-CALL.

    MOVE RETORNO TO WAL-MD5-RETORNO.

    CANCEL WA-LIB.

    .

RT00-10-FINAL.
    EXIT PROGRAM.
    STOP RUN.
