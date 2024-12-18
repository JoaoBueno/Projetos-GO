IDENTIFICATION DIVISION.
PROGRAM-ID.      GO-FCO.
*------------------------------------------------------------------------------*
* CHAMA A ROTINA(GO) QUE ACESSA O SERPRO E TRAZ OS DADOS DE CLIENTES (CNPJ)    *
*                                                                              *
* CRIACAO...: 14/12/2023 - BUENO                                               *
* ALTERACAO.:   /  /     -                                                     *
*                                                                              *
*------------------------------------------------------------------------------*
ENVIRONMENT DIVISION.
CONFIGURATION SECTION.
SPECIAL-NAMES.
    DECIMAL-POINT IS COMMA.

WORKING-STORAGE SECTION.
77  WN-LEN                      PIC 9(005)          VALUE ZEROS.

*LINKAGE SECTION.
77  LNK-FCO                     PIC 9(015)          VALUE  ZEROS.
77  LNK-RETORNO                 PIC X(1024)         VALUE SPACES.

*PROCEDURE DIVISION USING LNK-FCO LNK-RETORNO.
PROCEDURE DIVISION.
    SET CONFIGURATION "DLL-CONVENTION" TO 0.
    CALL "./go_serpro.so".

    MOVE 24907602000195 TO LNK-FCO.
    MOVE ALL X"00" TO LNK-RETORNO.
    MOVE LENGTH OF LNK-RETORNO TO WN-LEN.

    CALL "serpro"
             USING BY REFERENCE LNK-FCO, LNK-RETORNO
             BY REFERENCE WN-LEN
             GIVING RETURN-CODE
    END-CALL.

    DISPLAY LNK-FCO.
    DISPLAY LNK-RETORNO.

    ACCEPT LNK-FCO.

