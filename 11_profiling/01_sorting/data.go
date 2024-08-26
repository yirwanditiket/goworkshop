package main

var toSort []int = []int{
	943,
	815,
	6,
	469,
	51,
	863,
	199,
	102,
	430,
	21,
	171,
	144,
	718,
	952,
	806,
	183,
	555,
	249,
	394,
	624,
	350,
	234,
	818,
	61,
	882,
	860,
	241,
	492,
	881,
	28,
	971,
	509,
	6,
	19,
	115,
	253,
	491,
	92,
	251,
	980,
	403,
	878,
	124,
	717,
	153,
	732,
	970,
	247,
	965,
	263,
	858,
	414,
	826,
	716,
	138,
	510,
	915,
	238,
	936,
	711,
	917,
	137,
	664,
	849,
	669,
	871,
	470,
	528,
	154,
	214,
	442,
	128,
	902,
	884,
	60,
	981,
	536,
	288,
	204,
	330,
	294,
	506,
	552,
	942,
	280,
	795,
	606,
	971,
	936,
	277,
	912,
	766,
	358,
	196,
	675,
	63,
	439,
	105,
	801,
	225,
	60,
	528,
	526,
	826,
	996,
	108,
	147,
	934,
	855,
	348,
	15,
	654,
	561,
	856,
	998,
	171,
	557,
	453,
	611,
	548,
	521,
	575,
	703,
	757,
	479,
	646,
	736,
	220,
	122,
	171,
	82,
	996,
	906,
	408,
	496,
	377,
	586,
	495,
	491,
	666,
	948,
	122,
	590,
	535,
	393,
	482,
	399,
	47,
	20,
	453,
	17,
	654,
	744,
	221,
	505,
	615,
	849,
	832,
	887,
	632,
	656,
	768,
	365,
	654,
	266,
	697,
	234,
	813,
	526,
	244,
	798,
	553,
	772,
	552,
	62,
	959,
	506,
	565,
	614,
	633,
	492,
	559,
	827,
	768,
	229,
	684,
	179,
	779,
	771,
	134,
	305,
	767,
	433,
	606,
	771,
	329,
	191,
	24,
	66,
	858,
	675,
	937,
	221,
	795,
	796,
	962,
	2,
	917,
	918,
	464,
	211,
	899,
	391,
	960,
	746,
	927,
	315,
	702,
	122,
	28,
	366,
	575,
	553,
	373,
	403,
	946,
	401,
	336,
	117,
	718,
	388,
	682,
	1,
	452,
	523,
	284,
	771,
	584,
	756,
	515,
	547,
	857,
	524,
	667,
	439,
	53,
	983,
	582,
	808,
	712,
	868,
	88,
	146,
	171,
	512,
	672,
	593,
	204,
	972,
	75,
	316,
	703,
	734,
	769,
	20,
	583,
	864,
	230,
	133,
	456,
	266,
	626,
	944,
	779,
	78,
	791,
	271,
	303,
	405,
	760,
	589,
	222,
	85,
	34,
	721,
	519,
	799,
	927,
	936,
	298,
	776,
	993,
	63,
	350,
	850,
	983,
	226,
	880,
	982,
	616,
	504,
	536,
	719,
	35,
	208,
	509,
	322,
	849,
	78,
	612,
	711,
	234,
	659,
	884,
	279,
	903,
	188,
	576,
	426,
	816,
	414,
	107,
	623,
	108,
	965,
	267,
	207,
	237,
	312,
	579,
	516,
	239,
	288,
	671,
	87,
	567,
	49,
	361,
	990,
	286,
	880,
	696,
	536,
	13,
	229,
	33,
	335,
	536,
	128,
	8,
	405,
	92,
	598,
	510,
	413,
	755,
	314,
	114,
	816,
	712,
	127,
	396,
	26,
	545,
	128,
	569,
	174,
	621,
	780,
	704,
	646,
	969,
	134,
	582,
	356,
	5,
	667,
	815,
	304,
	434,
	312,
	113,
	205,
	582,
	475,
	474,
	652,
	32,
	395,
	839,
	426,
	441,
	458,
	559,
	829,
	173,
	428,
	129,
	717,
	283,
	912,
	275,
	60,
	922,
	682,
	663,
	833,
	154,
	264,
	865,
	132,
	816,
	890,
	355,
	191,
	441,
	730,
	16,
	335,
	472,
	197,
	123,
	803,
	296,
	452,
	827,
	345,
	150,
	402,
	841,
	158,
	639,
	850,
	150,
	419,
	605,
	322,
	255,
	540,
	705,
	860,
	404,
	242,
	734,
	717,
	82,
	202,
	690,
	435,
	78,
	810,
	890,
	369,
	313,
	862,
	35,
	987,
	188,
	694,
	51,
	449,
	881,
	334,
	51,
	10,
	422,
	793,
	234,
	829,
	174,
	393,
	134,
	255,
	118,
	973,
	46,
	8,
	631,
	272,
	442,
	132,
	59,
	749,
	258,
	362,
	466,
	729,
	756,
	627,
	37,
	989,
	957,
	511,
	325,
	497,
	798,
	369,
	367,
	737,
	983,
	357,
	251,
	417,
	828,
	977,
	295,
	115,
	937,
	491,
	742,
	275,
	904,
	244,
	297,
	371,
	417,
	363,
	776,
	607,
	305,
	904,
	50,
	374,
	386,
	572,
	661,
	639,
	541,
	319,
	456,
	471,
	750,
	500,
	943,
	301,
	557,
	832,
	782,
	284,
	242,
	784,
	518,
	834,
	934,
	62,
	19,
	404,
	156,
	181,
	26,
	629,
	905,
	690,
	502,
	164,
	604,
	582,
	189,
	149,
	795,
	758,
	303,
	133,
	238,
	609,
	657,
	778,
	810,
	453,
	1,
	242,
	119,
	875,
	120,
	261,
	520,
	390,
	187,
	106,
	181,
	550,
	490,
	137,
	946,
	170,
	801,
	597,
	35,
	4,
	482,
	987,
	291,
	272,
	135,
	388,
	65,
	833,
	543,
	908,
	478,
	166,
	535,
	584,
	98,
	229,
	181,
	329,
	667,
	21,
	698,
	699,
	786,
	183,
	665,
	943,
	60,
	628,
	705,
	75,
	975,
	887,
	705,
	170,
	117,
	205,
	687,
	266,
	812,
	338,
	175,
	828,
	333,
	188,
	433,
	91,
	882,
	397,
	637,
	386,
	467,
	904,
	131,
	349,
	903,
	283,
	208,
	84,
	295,
	991,
	566,
	579,
	748,
	129,
	966,
	943,
	593,
	300,
	792,
	503,
	780,
	22,
	292,
	489,
	861,
	690,
	294,
	592,
	785,
	441,
	565,
	303,
	924,
	504,
	805,
	127,
	496,
	419,
	496,
	700,
	790,
	994,
	585,
	616,
	318,
	824,
	744,
	744,
	497,
	302,
	96,
	111,
	921,
	542,
	644,
	354,
	519,
	747,
	819,
	691,
	821,
	625,
	701,
	485,
	316,
	360,
	164,
	91,
	459,
	494,
	45,
	279,
	349,
	620,
	67,
	379,
	89,
	180,
	432,
	712,
	826,
	606,
	816,
	346,
	185,
	371,
	88,
	428,
	587,
	416,
	176,
	551,
	198,
	123,
	206,
	632,
	181,
	736,
	609,
	516,
	976,
	152,
	306,
	990,
	940,
	801,
	973,
	253,
	664,
	749,
	770,
	536,
	784,
	2,
	715,
	668,
	520,
	251,
	334,
	518,
	756,
	283,
	215,
	880,
	416,
	692,
	354,
	935,
	135,
	347,
	795,
	350,
	910,
	476,
	377,
	780,
	144,
	683,
	50,
	555,
	951,
	912,
	723,
	734,
	647,
	910,
	469,
	663,
	740,
	830,
	740,
	236,
	244,
	102,
	681,
	395,
	560,
	602,
	664,
	629,
	675,
	731,
	117,
	743,
	945,
	752,
	957,
	808,
	458,
	709,
	570,
	384,
	378,
	963,
	644,
	756,
	722,
	837,
	395,
	869,
	444,
	949,
	552,
	480,
	16,
	845,
	310,
	36,
	993,
	229,
	905,
	524,
	360,
	103,
	205,
	688,
	151,
	337,
	337,
	146,
	647,
	290,
	942,
	171,
	247,
	322,
	616,
	634,
	219,
	444,
	228,
	630,
	400,
	514,
	318,
	177,
	212,
	400,
	351,
	684,
	349,
	118,
	597,
	216,
	480,
	128,
	451,
	688,
	203,
	544,
	54,
	31,
	49,
	610,
	467,
	938,
	110,
	423,
	791,
	946,
	229,
	945,
	452,
	953,
	667,
	613,
	881,
	865,
	914,
	550,
	630,
	134,
	957,
	42,
	537,
	891,
	43,
	163,
	208,
	417,
	245,
	324,
	137,
	473,
	475,
	838,
	309,
	245,
	172,
	709,
	906,
	350,
	271,
	493,
	898,
	954,
	984,
	525,
	474,
	390,
	718,
	874,
	844,
	164,
	52,
	953,
	723,
	907,
	412,
	156,
	212,
	903,
	744,
	272,
	915,
	349,
	521,
	143,
	639,
	974,
	876,
	965,
	159,
	465,
	118,
	778,
	397,
	674,
	284,
	760,
	80,
	793,
	73,
	816,
	520,
	876,
	438,
	286,
	56,
	300,
	496,
	727,
	583,
	516,
	516,
	181,
	247,
	915,
	209,
	417,
	107,
	298,
	896,
	106,
	789,
	828,
	388,
	64,
	549,
	813,
	159,
	568,
	916,
	363,
	144,
	468,
	277,
	984,
	63,
	914,
	863,
	629,
	836,
	623,
	540,
	191,
	678,
	879,
	773,
	850,
	94,
}