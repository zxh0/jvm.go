package exception;

public class CatchTest {
    
    public static void main(String[] args) {
        if (f0() != 0) {
            System.out.println("f0() failed!");
        }
        if (f1() != 1) {
            System.out.println("f1() failed!");
        }
        if (f2() != 2) {
            System.out.println("f2() failed!");
        }
        if (f3() != 3) {
            System.out.println("f3() failed!");
        }
        System.out.println("OK!");
    }
    
    private static int f0() {
        try {
            bad();
            return -1;
        } catch (Throwable t) {
            return 0;
        }
    }
    
    private static int f1() {
        try {
            bad();
            return -1;
        } catch (Exception e) {
            return 1;
        }
    }
    
    private static int f2() {
        try {
            bad();
            return -1;
        } catch (RuntimeException e) {
            return 2;
        }
    }
    
    private static int f3() {
        try {
            bad2();
            return -1;
        } catch (RuntimeException e) {
            return 3;
        }
    }
    
    private static void bad() {
        throw new RuntimeException("BAD!");
    }
    
    private static void bad2() {
        try {
            bad();
        } catch (RuntimeException e) {
            throw e;
        }
    }
    
}
