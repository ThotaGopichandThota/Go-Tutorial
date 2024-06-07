package main

import (
	"fmt"
)

func main() {
	//var name string
	var name string = "my name is gopichand"
	// fmt.Println("enter the string :")
	// fmt.Scanln(&name)
	fmt.Println("the string is:", name)
	fmt.Println("length of the string:", len(name))
	for i := 0; i < len(name); i++ {
		//	fmt.Println("the characters present int string:", string(name[i]))
		char := string(name[i])
		fmt.Println(char)
	}

}
// public class StringRotation {

// 	public static void main(String[] args) 
// 	{
// 	 Scanner sc =new Scanner(System.in);
//       System.out.println("enter string:");
//       String s=sc.next();
//       System.out.println("enter no of rotations");
//       int n=sc.nextInt();
//      System.out.println( rotateString(s, n));
// 	}
//    public static String rotateString(String s,int n)
//    {
// 	   String rotated=s.substring(n)+s.substring(0, n);
// 	   String[]words=rotated.split(s);
// 	   StringBuilder res=new StringBuilder();
// 	   for(String word:words)
// 	   {
// 		   res.append(word).append(" ");
// 	   }
// 	   return res.toString().trim();
//    }