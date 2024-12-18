function welcome()	{
//	alert("Hi!");
}

function checkNumber(input, min, max, msg)
{
    msg = msg + " field has invalid data: " + input;
    var str = input;
    for (var i = 0; i < str.length; i++) {
        var ch = str.substring(i, i + 1)
        if ((ch < "0" || "9" < ch) && ch != '.') {
            alert(msg);
            return false;
        }
    }
    var num = parseFloat(str)
    if (num < min || max <= num) {
        alert(msg + " not in range [" + min + ".." + max + "]");
        return false;
    }
    input = str;
    return true;
}

function computeField(input)
{
  if (input != null && input.length != 0)
	input = "" + eval(input);
  computeForm(input.form);
}

function computeInverseField(input)
{
  if (input != null && input.length != 0)	{
		input = "" + eval(input);
	  computeInverseForm(input.form);
	}
}

function MakeArray(size)
{
	this.length = size;
  for(var i = 1; i <= size; i++)
	{
		this[i] = 0;
	}
	return this;
}

function monthLength(month, leap)
{
	monthLengthArray = new MakeArray(12);

	monthLengthArray[1] = 32;
	monthLengthArray[2] = 29 + leap;
	monthLengthArray[3] = 32;
	monthLengthArray[4] = 31;
	monthLengthArray[5] = 32;
	monthLengthArray[6] = 31;
	monthLengthArray[7] = 32;
	monthLengthArray[8] = 32;
	monthLengthArray[9] = 31;
	monthLengthArray[10] = 32;
	monthLengthArray[11] = 31;
	monthLengthArray[12] = 32;

	return monthLengthArray[month];
}

function computeTheForm(month, day, year, julianDay)
{
	var leap = (
		(gregorian == false) && (year%4 == 0)? 1:
		(year%4 != 0? 0: 
		(year%400 == 0? 1:
		(year%100==0? 0:
		1))));
    if (!checkNumber(month, 1, 13, "Month") ||
        !checkNumber(day, 1, monthLength(parseFloat(month), leap), "Day") ||
        !checkNumber(year, 1,1000000, "Year")) {
        julianDay = "Invalid";
        return;
    }
	var D = eval(day);
	var M = eval(month);
	var Y = eval(year);
	if(M<3)	{
		Y--;
		M += 12;
	}

//alert("D= " + D);
//alert("M= " + M);
//alert("Y= " + Y);
	if(gregorian == true)	{
		var A = Math.floor(Y/100);
		var B = Math.floor(A/4);
		var C = 2 - A + B;
	}
	else
		C=0;
//alert("C= " + C);
	var E = Math.floor(365.25*(Y + 4716));
//alert("E= " + E);
	var F = Math.floor(30.6001*(M + 1));
//alert("F= " + F);
	var julianday = C + eval(D) + E + F - 1524.5;
//	if(julianday<2299160.5)	{
//		alert("The date you have entered is before the introduction of the Gregorian calendar!");
//		NewJD = "Invalid";
//	}
//	else
 	 NewJD = julianday;

//	if(julianday - Math.floor(julianday) != 0.5)
//	 NewJD = "Invalid";

	if(julianDay == null || julianDay == 0)
		julianDay = NewJD;
	else
		if(julianDay != NewJD) julianDay = NewJD;
}

function computeInverseForm(month, day, year, julianDay)
{
	if(!checkNumber(julianDay,2299160.5,100000000000,"Julian Day Number"))	{
		alert("Date entered before October 15, 1582. Result is on Gregorian Proleptic Calendar.");
//		julianDay = "Invalid";
//		return;
	}
	var JD = eval(julianDay);

	Z = JD+0.5;
	F = Z - Math.floor(Z);

	if(gregorian == true)	{
		Z = Math.floor(Z);
		W = Math.floor((Z - 1867216.25)/36524.25);
		X = Math.floor(W/4);
		A = Z + 1 + W - X;
	}
	else
		A = Z;
	B = A + 1524;
	C = Math.floor((B - 122.1)/365.25);
	D = Math.floor(365.25*C);
	E = Math.floor((B - D)/30.6001);

	NewMonth = E>13? E-13: E-1;
	NewDay = B - D - Math.floor(30.6001*E) +F;
	NewYear = NewMonth<3? C-4715: C-4716;

  
    if ((month == null || month == 0) ||
        (day == null || day == 0) ||
        (year == null || year == 0)) {
		month = NewMonth;
		day = NewDay;
		year = NewYear;
	}
	else {
		if(month != NewMonth) month = NewMonth;
		if(day != NewDay) day = NewDay;
		if(year != NewYear) year = NewYear;
	}
}

function computeForm(month, day, year, julianDay)
{
    if ((month == null || month == 0) ||
        (day == null || day == 0) ||
        (year == null || year == 0)) 
             {
             if(julianDay == null || julianDay == 0)
                  return;
             else
                  computeInverseForm(month, day, year, julianDay);
    }
	computeTheForm(month, day, year, julianDay);
}

function clearForm(form)
{
    month = "";
    day = "";
    year = "";
    julianDay = "";
//	alert(document.forms[0].calendar[0].checked);
//	alert(document.forms[0].calendar[1].checked);
}

month = 0;
day = 0;
year = 0;
julianDay = 2456293;
gregorian = true;
computeForm(month, day, year, julianDay);

