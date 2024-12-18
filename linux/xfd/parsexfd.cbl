identification division.
program-id.	ParseXFD.
author.		Randy Zack, Acucorp, Inc.
		This program Parses and returns info about XFD files.

	RCS INFO: $Id: ParseXFD.cbl 50338 2007-02-09 16:47:51Z rzack $

data division.
working-storage section.

copy "bacu/parsexfd.ws".

77  parsexfd-program-version	pic x(5) value " 2.0".

01  xfd-pointer			pointer value null.

01  save-col-case		pic 9.
01  save-xfd-index-for-cond     unsigned-short.
01  temp-idx                    pic 99.
01  temp-idx2                   pic 99.
01  temp-values.
    03  value-9                 pic 9.
    03  value-99                pic 99.
    03  value-999               pic 999.
    03  value-9999              pic 9999.
    03  value-99999             pic 99999.

linkage section.
77  opcode			pic 9.

procedure division using opcode.
main-level.
    evaluate opcode
      when parse-xfd-op
	if xfd-pointer not = null
	    move XFDParseXFDAlreadyParsedError to parse-flag
	else
	    perform parse-xfd
	end-if

      when get-key-info-op
	evaluate true
	  when xfd-key-index > xfd-number-of-keys
	    move XFDParseInvalidKeyIndexError to parse-flag
	  when xfd-pointer = null
	    move XFDParseNoXFDParsedError to parse-flag
	  when other
	    perform get-key-info
	end-evaluate

      when get-cond-info-op
	evaluate true
	  when xfd-cond-index > xfd-number-conditions
	    move XFDParseInvalidCondIndexError to parse-flag
	  when xfd-pointer = null
	    move XFDParseNoXFDParsedError to parse-flag
	  when other
	    perform get-cond-info
	end-evaluate

      when get-field-info-op
	evaluate true
	  when xfd-field-index > xfd-total-number-fields
	    move XFDParseInvalidFieldIndexError to parse-flag
	  when xfd-pointer = null
	    move XFDParseNoXFDParsedError to parse-flag
	  when other
	    perform get-field-info
	end-evaluate

      when test-conditions-op
	if xfd-pointer = null
	    move XFDParseNoXFDParsedError to parse-flag
	else
	    perform test-all-conditions
	end-if

      when free-memory-op
	if xfd-pointer not = null
	    call "C$PARSEXFD" using PARSEXFD-RELEASE, xfd-pointer
	end-if
	set xfd-pointer to null
    end-evaluate.
    exit program.

***   The variables are 1-based, but the indexes internally are 0-based.
get-key-info.
    subtract 1 from xfd-key-index.
    call "C$PARSEXFD" using PARSEXFD-GET-KEY-INFO, xfd-pointer,
		xfd-key-index, PARSEXFD-KEY-DESCRIPTION.
    add 1 to xfd-key-index.
    if return-code = 0
	move XFDParseInvalidKeyIndexError to parse-flag
    else
	move parsexfd-number-segments to xfd-number-of-segments
	set xfd-allow-duplicates to false
	if parsexfd-allow-duplicates
	    set xfd-allow-duplicates to true
	end-if
	perform varying parsexfd-seg-idx from 1 by 1 until
		parsexfd-seg-idx > parsexfd-number-segments
	    move parsexfd-segment-length(parsexfd-seg-idx) to
			xfd-segment-length(parsexfd-seg-idx)
	    move parsexfd-segment-offset(parsexfd-seg-idx) to
			xfd-segment-offset(parsexfd-seg-idx)
	end-perform
	move parsexfd-number-key-fields to xfd-num-of-key-fields
	perform varying parsexfd-key-field-idx from 1 by 1 until
		parsexfd-key-field-idx > parsexfd-number-key-fields
	    add 1 to parsexfd-key-field-num(parsexfd-key-field-idx) giving
			xfd-key-field-num(parsexfd-key-field-idx)
	    call "C$PARSEXFD" using PARSEXFD-GET-FIELD-INFO, xfd-pointer,
		parsexfd-key-field-num(parsexfd-key-field-idx),
		PARSEXFD-FIELD-DESCRIPTION
	    move parsexfd-field-name to
			xfd-key-field-name(parsexfd-key-field-idx)
	end-perform
    end-if.

get-cond-info.
    subtract 1 from xfd-cond-index.
    call "C$PARSEXFD" using PARSEXFD-GET-COND-INFO, xfd-pointer,
		xfd-cond-index, PARSEXFD-CONDITION-DESCRIPTION.
    add 1 to xfd-cond-index.
    if return-code = 0
	move XFDParseInvalidCondIndexError to parse-flag
    else
	move parsexfd-condition-type to xfd-condition-type
	set true-condition to false
	if parsexfd-true-condition
	    set true-condition to true
	end-if
	if parsexfd-comparison-cond or parsexfd-other-condition
	    move parsexfd-comp-fieldnum to xfd-other-fieldnum
	    move parsexfd-comp-fieldname to xfd-other-fieldname
	    if not parsexfd-other-condition
		move parsexfd-comp-field-val to xfd-other-field-val
	    end-if
	end-if
	if parsexfd-and-condition or parsexfd-or-condition
	    move parsexfd-condition-1 to xfd-condition-1
	    move parsexfd-condition-2 to xfd-condition-2
	end-if
	move parsexfd-condition-tablename to xfd-condition-tablename
    end-if.

get-field-info.
    subtract 1 from xfd-field-index.
    call "C$PARSEXFD" using PARSEXFD-GET-FIELD-INFO, xfd-pointer,
		xfd-field-index, PARSEXFD-FIELD-DESCRIPTION.
    add 1 to xfd-field-index.
    if return-code = 0
	move XFDParseInvalidFieldIndexError to parse-flag
    else
	move parsexfd-field-offset to xfd-field-offset
	move parsexfd-field-length to xfd-field-length
	move parsexfd-field-type to xfd-field-type
	move parsexfd-field-digits to xfd-field-digits
	move parsexfd-field-scale to xfd-field-scale
	move parsexfd-field-user-type to xfd-field-user-type
	move parsexfd-field-condition to xfd-field-condition
	move parsexfd-field-level to xfd-field-level
	move parsexfd-field-name to xfd-field-name
	move parsexfd-field-format to xfd-field-format
	move parsexfd-field-occurs-depth to xfd-field-occurs-depth
	perform varying parsexfd-field-occurs-level from 1 by 1
		until parsexfd-field-occurs-level > parsexfd-field-occurs-depth
	    move 0 to temp-idx
	    inspect xfd-field-name tallying temp-idx for trailing spaces
	    set temp-idx2 to size of xfd-field-name
	    compute temp-idx = temp-idx2 - temp-idx + 1
	    if parsexfd-field-occurs-level = 1
		move "(" to xfd-field-name(temp-idx:)
	    else
		move "," to xfd-field-name(temp-idx:)
	    end-if
	    add 1 to temp-idx
	    if parsexfd-field-occ-this-idx(parsexfd-field-occurs-level) < 10
		move parsexfd-field-occ-this-idx(parsexfd-field-occurs-level)
			to value-9
		move value-9 to xfd-field-name(temp-idx:)
		add 1 to temp-idx
	    else if parsexfd-field-occ-this-idx(parsexfd-field-occurs-level) < 100
		move parsexfd-field-occ-this-idx(parsexfd-field-occurs-level)
			to value-99
		move value-99 to xfd-field-name(temp-idx:)
		add 2 to temp-idx
	    else if parsexfd-field-occ-this-idx(parsexfd-field-occurs-level) < 1000
		move parsexfd-field-occ-this-idx(parsexfd-field-occurs-level)
			to value-999
		move value-999 to xfd-field-name(temp-idx:)
		add 3 to temp-idx
	    else if parsexfd-field-occ-this-idx(parsexfd-field-occurs-level) < 10000
		move parsexfd-field-occ-this-idx(parsexfd-field-occurs-level)
			to value-9999
		move value-9999 to xfd-field-name(temp-idx:)
		add 4 to temp-idx
	    else
		move parsexfd-field-occ-this-idx(parsexfd-field-occurs-level)
			to xfd-field-name(temp-idx:)
		add 5 to temp-idx
	    end-if end-if end-if end-if
	    move parsexfd-field-occ-max-idx(parsexfd-field-occurs-level)
			to xfd-field-occ-max-idx(parsexfd-field-occurs-level)
	    if parsexfd-field-occurs-level = parsexfd-field-occurs-depth
		move ")" to xfd-field-name(temp-idx:)
	    end-if
	end-perform
    end-if.

*  Test conditions.  This assumes that all conditions are only dependent on
*  earlier-numbered conditions.  This assumption has been verified.
test-all-conditions.
    call "C$PARSEXFD" using PARSEXFD-TEST-CONDITIONS, xfd-pointer,
		record-area-ptr.
*  Find the min and max valid fields.
    move 0 to max-xfd-field-index.
    add 1 to xfd-total-number-fields giving min-xfd-field-index.
    move xfd-field-index to save-xfd-index-for-cond.
    perform varying xfd-field-index from 1 by 1 until
                        xfd-field-index > xfd-total-number-fields
        perform get-field-info
        if xfd-field-condition not = 0
            move xfd-field-condition to xfd-cond-index
            perform get-cond-info
        end-if
        if xfd-field-condition = 0 or true-condition
            if xfd-field-index < min-xfd-field-index
                move xfd-field-index to min-xfd-field-index
            end-if
            if xfd-field-index > max-xfd-field-index
                move xfd-field-index to max-xfd-field-index
            end-if
        end-if
    end-perform.
    move save-xfd-index-for-cond to xfd-field-index.


*** The previous version of alfred (ParseXFD) didn't modify the field names
*** as they exist in the XFD.  So we need to set 4GL-COL-CASE to UNCHANGED.
*** But to be really nice, we should save the old value, and restore it
*** when finished.
parse-xfd.
    accept save-col-case from environment "4GL-COLUMN-CASE".
    set environment "4GL-COLUMN-CASE" to "UNCHANGED".
    call "C$PARSEXFD" using PARSEXFD-PARSE, xfdfile, filename,
		PARSEXFD-FLAG-DEEP-FIRST, parsexfd-description.
    set environment "4GL-COLUMN-CASE" to save-col-case.
    move return-unsigned to xfd-pointer.
    if xfd-pointer = null
***   Set the error code based on f-errno and f-int-errno
	evaluate true
	  when E-NO-MEMORY
	    move XFDParseNoMemoryError to parse-flag
	  when E-MISMATCH
	    move XFDParseMismatchError to parse-flag
	  when E-INTERFACE
	    evaluate f-int-errno
	      when 1
		move XFDParseReadError to parse-flag
	      when 2
		move XFDParseVersionError to parse-flag
	      when 3
		move XFDParseOpenError to parse-flag
	      when 4
		move XFDParseTooManyKeyFieldsError to parse-flag
	      when 6
		move XFDParseMismatchError to parse-flag
	    end-evaluate
	end-evaluate
	exit paragraph
    end-if.

* Allow only indexed files at this point
    if parsexfd-filetype not = 12
	move XFDParseInvalidFileType to parse-flag
	call "C$PARSEXFD" using PARSEXFD-RELEASE, xfd-pointer
	exit paragraph
    end-if.

***  We no longer need to perform a reality check - C$PARSEXFD does
***  that for us!
    move parsexfd-version to xfd-version.
    move parsexfd-select-name to xfd-select-name.
    move parsexfd-filename to xfd-filename.
    move parsexfd-filetype to xfd-filetype.
    move parsexfd-max-rec-size to xfd-max-record-size.
    move parsexfd-min-rec-size to xfd-min-record-size.
    move parsexfd-num-keys to xfd-number-of-keys.
    move parsexfd-number-conditions to xfd-number-conditions.
    move parsexfd-number-fields to xfd-number-fields xfd-total-number-fields.
    move 1 to xfd-field-index, xfd-key-index.

    perform varying xfd-key-index from 0 by 1
		until xfd-key-index = parsexfd-num-keys
	call "C$PARSEXFD" using PARSEXFD-GET-KEY-INFO, xfd-pointer,
		xfd-key-index, PARSEXFD-KEY-DESCRIPTION
	move parsexfd-number-key-fields to xfd-num-key-flds(xfd-key-index + 1)
    end-perform.
