*  Working-Storage section definitions for programs that use ParseXFD.

***   The following file sets up many of the constants that used to be here.
copy "bacu/parsexfd.def".

***   First, lots of level-78 values

* Parse Errors - Start at 2 so this just indexes into the message file
* If you are using ParseXFD, you won't have a message file
78  XFDParseOpenError		    value 2.
78  XFDParseReadError		    value 3.
78  XFDParseVersionError	    value 4.
78  XFDParseMismatchError	    value 5.
78  XFDParseSizeError		    value 6.
78  XFDParseNoMemoryError	    value 7.
78  XFDParseTooManyKeyFieldsError   value 8.
78  XFDParseNoXFDParsedError	    value 9.
78  XFDParseXFDAlreadyParsedError   value 10.
78  XFDParseInvalidKeyIndexError    value 11.
78  XFDParseInvalidCondIndexError   value 12.
78  XFDParseInvalidFieldIndexError  value 13.
78  XFDParseInvalidFileType         value 14.

* Set the record-area-ptr to the address of you record area.
* Set the filename to the name of the file described by the XFD.
* Set the xfdfile to the name of the XFD file.
* After any ParseXFD operation, parse-flag will be set to any error.
01  record-area-ptr		pointer sync	external.
01  filename			pic x(78)	external.
01  xfdfile			pic x(78)	external.
01  ext-flags					external.
    03  parse-flag		pic 99.
	88  parse-error			values 1 thru 99 false 0.

* Describe one key from the XFD file.  This group-item should be
* considered as READ-ONLY!!
01  xfd-key-description				external.
    03  xfd-number-of-segments	pic 99.
    03  xfd-allow-dup-flag	pic 9.
	88  xfd-allow-duplicates	value 1 false 0.
    03  xfd-segment-description	occurs max-segs times indexed by xfd-seg-idx.
	05  xfd-segment-length	pic x	comp-x.
	05  xfd-segment-offset	pic x(4)	comp-x.
    03  xfd-num-of-key-fields	pic x	comp-x.
    03  xfd-key-fields		occurs MaxNumKeyFields times
				indexed by xfd-key-field-idx.
	05  xfd-key-field-name	pic x(30).
	05  xfd-key-field-num	pic xx	comp-x.

* Describe one condition from the XFD file.  This group-item should be
* considered as READ-ONLY!!
01  xfd-condition-description			external.
    03  xfd-condition-type	pic x	comp-x.
	88  xfd-equal-condition		value 1.
	88  xfd-and-condition		value 2.
	88  xfd-other-condition		value 3.
	88  xfd-gt-condition		value 4.
	88  xfd-ge-condition		value 5.
	88  xfd-lt-condition		value 6.
	88  xfd-le-condition		value 7.
	88  xfd-ne-condition		value 8.
	88  xfd-or-condition		value 9.
	88  xfd-comparison-condition	values 1, 4 through 8.
    03  xfd-condition-flag	pic x.
	88  true-condition		value 'y' false 'n'.
    03  xfd-other-conditions.
	05  xfd-other-fieldnum	pic xx	comp-x.
	05  xfd-other-fieldname	pic x(30).
	05  xfd-other-field-val	pic x(50).
	05  xfd-other-field-nums redefines xfd-other-field-val.
	    07  xfd-cond-val-1	pic s9(18).
	    07  xfd-cond-val-1V99 redefines xfd-cond-val-1
			        pic s9(16)v99.
	    07  xfd-cond-val-2	pic s9(18).
    03  xfd-and-conditions	redefines xfd-other-conditions.
	05  xfd-condition-1	pic xx	comp-x.
	05  xfd-condition-2	pic xx	comp-x.
    03  xfd-condition-tablename	pic x(30).

* Describe one field from the XFD file.  This group-item should be
* considered READ-ONLY!!
01  xfd-field-description			external.
    03  xfd-field-offset	pic x(4)	comp-x.
    03  xfd-field-length	pic x(4)	comp-x.
    03  xfd-field-type		pic x		comp-x.
        88  signed-field		values	NumSignSep
                                                NumSigned
                                                NumSepLead
                                                NumLeading
                                                CompSigned
                                                PackedSigned
                                                BinarySigned
                                                NativeSigned.
        88  numeric-field		values	NumEdited thru NativeUnsigned.
        88  float-field			value	Flt.
        88  ascii-field			values	Alphanum thru Group.
        88  national-field		values	Nat-type thru NatEdited.
        88  wide-field			values	Wide-type thru WideEdited.
    03  xfd-field-digits	pic x(4)	comp-x.
    03  xfd-field-scale		pic s99	comp-4.
    03  xfd-field-user-type	pic xx	comp-x.
    03  xfd-field-condition	pic xx	comp-x.
    03  xfd-field-level		pic x	comp-x.
    03  xfd-field-name		pic x(30).
    03  xfd-field-format        pic x(30).
    03  xfd-field-occurs-depth	pic x	comp-x.
    03  xfd-field-occurs-table	occurs MaxNumKeyFields times
				indexed by xfd-field-occurs-level.
	05  xfd-field-occ-max-idx pic xx comp-x.
	05  xfd-field-occ-offset pic xx	comp-x.

* The entire XFD description.  Most of these fields are READ-ONLY.
* You need to modify xfd-key-index, xfd-field-index and xfd-cond-index
* in order to get information about a particular key, field or condition.
* But don't modify anything else.
01  xfd-description				external.
    03  xfd-version		pic x	comp-x.
    03  xfd-select-name		pic x(30).
    03  xfd-filename		pic x(30).
    03  xfd-filetype		pic x   comp-x.
    03  xfd-max-record-size	pic x(4)	comp-x.
    03  xfd-min-record-size	pic x(4)	comp-x.
    03  xfd-number-of-keys	pic x	comp-x.
    03  xfd-number-conditions	pic xx	comp-x.
    03  xfd-number-fields	pic xx	comp-x.
    03  xfd-total-number-fields	pic xx	comp-x.
    03  xfd-key-index		pic xx	comp-x.
    03  xfd-field-index		pic xx	comp-x.
    03  save-xfd-field-index	pic xx	comp-x.
    03  min-xfd-field-index	pic xx	comp-x.
    03  max-xfd-field-index	pic xx	comp-x.
    03  xfd-cond-index		pic xx	comp-x.
    03  xfd-max-field-name-len	pic xx	comp-x.
    03  xfd-num-key-flds	pic x	comp-x	occurs 120 times.

* Valid operations for using when calling ParseXFD
78  parse-xfd-op		value 0.
78  get-key-info-op		value 1.
78  get-cond-info-op		value 2.
78  get-field-info-op		value 3.
78  test-conditions-op		value 4.
78  free-memory-op		value 9.
