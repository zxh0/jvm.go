public class StringTest {
    
    public static void main(String[] args) {
        String s1 = "abc1";
        String s2 = "abc1";
        if (s1 == s2) {
            System.out.println("OK1!");
        } else {
            System.out.println("Fail1!");
        }
        
        int x = 1;
        String s3 = "abc" + x;
        if (s1 != s3) {
            System.out.println("OK2!");
        } else {
            System.out.println("Fail2!");
        }
        
        s3 = s3.intern();
        if (s1 == s3) {
            System.out.println("OK3!");
        } else {
            System.out.println("Fail3!");
        }
    }
    
}
