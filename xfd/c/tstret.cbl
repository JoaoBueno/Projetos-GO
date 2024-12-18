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
77  WA-LIB                      PIC X(150)              VALUE SPACES.
77  STR                         PIC X(100)              VALUE SPACES.
77                              PIC X(001)              VALUE X"00".
77  RETORNO                     PIC X(032)              VALUE SPACES.
77                              PIC X(001)              VALUE X"00".
77  RETORNO1                    PIC X(032)              VALUE SPACES.
77                              PIC X(001)              VALUE X"00".
77  BRUNO                       PIC X(100)              VALUE
    "123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890".
77  LEN                         SIGNED-LONG.
77                              PIC X(001)              VALUE X"00".


PROCEDURE DIVISION.
RT00-00-INICIO.
* 0 - FUNCOES EM C
    SET CONFIGURATION "DLL-CONVENTION" TO 0.

*    CALL "./t1c.so".
*    MOVE "teste" TO STR.
* 
*    INSPECT STR TALLYING LEN FOR CHARACTERS BEFORE INITIAL "  ".
* 
*    CALL "retornamd5"
*             USING BY REFERENCE STR, RETORNO
*             BY REFERENCE LEN
*             GIVING RETURN-CODE
*    END-CALL.

    DISPLAY BRUNO AT 0101.
    CALL "./t2c.so".

    MOVE "teste" TO STR.
    MOVE ALL X"00" TO RETORNO.
    MOVE LENGTH OF RETORNO TO LEN.

    CALL "retornamd5"
             USING BY REFERENCE STR, RETORNO
             BY REFERENCE LEN
             returning RETORNO1
    END-CALL.

    ACCEPT WAP-USAMD5.
