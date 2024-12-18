PROGRAM-ID.  PIDTOUSER.

FILE-CONTROL.
	SELECT PS-FILE
	ASSIGN TO PS-COMMAND
	LINE SEQUENTIAL.

FILE SECTION.
FD  PS-FILE.
01  PS-RECORD                   PIC X(080).

WORKING-STORAGE SECTION.
77  teste          pic 9(005).
77  FILLER         PIC X VALUE X"00".
77  MY-SHARED-INT  SIGNED-INT.
77  FILLER         PIC X VALUE X"00".
77  RETORNO        SIGNED-INT.
77  FILLER         PIC X VALUE X"00".
77  RETORNA        PIC X(080).
77  FILLER         PIC X VALUE X"00".
77  NOME           PIC X(080).
77  FILLER         PIC X VALUE X"00".
* ps hp 29779 -o user
01  PS-COMMAND.
    03                          PIC X(009) VALUE "-P ps hp ".
    03  COMMAND-PID             PIC Z(005).
    03                          PIC X(008) VALUE " -o user".

*LINKAGE SECTION.
77  PID                         PIC 9(005).
77  USER-NAME                   PIC X(030).

*PROCEDURE DIVISION CHAINING PID.
PROCEDURE DIVISION CHAINING PID.
MAIN-LOGIC.
    SET CONFIGURATION "DLL-CONVENTION" TO 0.
    initialize NOME RETORNA.
    MOVE "select * from cliente" TO NOME.
    call "./helloworld.so".
    call "foo" using BY REFERENCE NOME, RETORNA GIVING teste.

    DISPLAY NOME.
    DISPLAY RETORNA.
    DISPLAY teste.
    DISPLAY "TCHAU".


    *> MOVE PID TO COMMAND-PID.
    *> OPEN INPUT PS-FILE.
    *> READ PS-FILE NEXT
    *>      RECORD INTO USER-NAME
    *>      AT END MOVE SPACES TO USER-NAME
    *> END-READ.

    *> DISPLAY USER-NAME upon sysout.
    *> CLOSE PS-FILE.
    *> EXIT PROGRAM.
