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
77  WA-PARA                     PIC X(001)          VALUE SPACES.

01  WA-REGIS.
    03  WN-FCO                  PIC 9(014)         VALUE ZEROS.
    03  WA-NOME                 PIC X(055)         VALUE SPACES.
    03  WA-JURFIS               PIC X(001)         VALUE SPACES.
    03  WN-CNPJ                 PIC 9(014)         VALUE ZEROS.
    03  WA-CIDADE               PIC X(050)         VALUE SPACES.
    03  WA-UF                   PIC X(002)         VALUE SPACES.
77  WN-RET                      SIGNED-LONG.

*LINKAGE SECTION.
77  LNK-FCO                     PIC 9(014)          VALUE  ZEROS.
77  LNK-RETORNO                 PIC X(1024)         VALUE SPACES.

*PROCEDURE DIVISION USING LNK-FCO LNK-RETORNO.
PROCEDURE DIVISION.
    SET CONFIGURATION "DLL-CONVENTION" TO 0.
    CALL "./go_bd.so".

    MOVE SPACES TO WA-PARA.

    CALL "fco_bd"
             USING BY REFERENCE LNK-FCO, LNK-RETORNO, WN-RET
             returning INTO WN-RET
    END-CALL.

    if WN-RET < 0
        DISPLAY "ERRO: " WN-RET
        ACCEPT LNK-FCO
        STOP RUN.    

    IF   LNK-RETORNO(1:9) = "<<<fim>>>"
         MOVE "S" TO WA-PARA
    END-IF.

    MOVE LNK-RETORNO TO WA-REGIS.
    DISPLAY MESSAGE BOX "FCO: " WN-FCO X"0A"
                        "NOME: " WA-NOME X"0A"
                        "CNPJ: " WN-CNPJ X"0A"
                        "CIDA: " WA-CIDADE X"0A"
                        "UF: " WA-UF.

    PERFORM UNTIL WA-PARA = "S"
            CALL "fco_next"
                     USING BY REFERENCE LNK-FCO, LNK-RETORNO
                     BY REFERENCE WN-RET
                     returning INTO WN-RET
            END-CALL
            IF   LNK-RETORNO(1:9) = "<<<fim>>>"
                 MOVE "S" TO WA-PARA
            ELSE MOVE LNK-RETORNO TO WA-REGIS
                 DISPLAY WA-REGIS
            END-IF
    END-PERFORM.

    CALL "closeRows".

    ACCEPT LNK-FCO.
