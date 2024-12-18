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
77  NomeEmpresarial     pic x(60).
77  NomeFantasia     pic x(60).
77  DataAbertura     pic x(60).
77  CorreioEletronico     pic x(60).
77  Porte     pic x(60).
77  EnderecoTipoLogradouro     pic x(60).
77  EnderecoLogradouro     pic x(60).
77  EnderecoBairro     pic x(60).
77  EnderecoMunicipioDescricao     pic x(60).
77  EnderecoUF     pic x(60).
77  EnderecoCEP     pic x(60).
77  Telefones1    pic x(60).
77  Telefones2     pic x(60).
77  Telefones3     pic x(60).
77  Telefones4     pic x(60).
77  CnaePrincipal     pic x(60).
77  CnaeSecundarias     pic x(60).

77  WN-RET                      SIGNED-LONG.

*LINKAGE SECTION.
77  LNK-FCO                     PIC 9(014)          VALUE  ZEROS.
77  LNK-RETORNO                 PIC X(1024)         VALUE SPACES.

*PROCEDURE DIVISION USING LNK-FCO LNK-RETORNO.
PROCEDURE DIVISION.
    SET CONFIGURATION "DLL-CONVENTION" TO 0.
    CALL "./go_serpro.so".

    MOVE 24907602000195 TO LNK-FCO.
    MOVE ALL X"00" TO LNK-RETORNO.
    MOVE LENGTH OF LNK-RETORNO TO WN-RET.

    CALL "serpro"
             USING BY REFERENCE LNK-FCO, LNK-RETORNO
             BY REFERENCE WN-RET
             returning INTO WN-RET
    END-CALL.

    DISPLAY LNK-FCO.
    DISPLAY LNK-RETORNO.
    
    unstring LNK-RETORNO delimited by "|" into
        NomeEmpresarial
        NomeFantasia
        DataAbertura
        CorreioEletronico
        Porte
        EnderecoTipoLogradouro
        EnderecoLogradouro
        EnderecoBairro
        EnderecoMunicipioDescricao
        EnderecoUF
        EnderecoCEP
        Telefones1
        Telefones2
        Telefones3
        Telefones4
        CnaePrincipal
        CnaeSecundarias.


    ACCEPT LNK-FCO.

