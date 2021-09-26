package main

func Temperature(A,B,C,D,E int) int  {
	if A < 0 {
		if B > 0 {
			return C*(-A)+D+B*E
		}else{
			return (B-A)*C
		}
	}else {
		return (B-A)*E
	}
}
