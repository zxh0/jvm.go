package exception.jvm;

public class NPETest {
    
    public static void main(String[] args) {
        try {
            arraylength(null);
        } catch (NullPointerException e) {
            e.printStackTrace(System.err);
        }
    }
    
    private static void arraylength(int[] x) {
        if (x.length > 0) {}
    }
    
    
//    private static void test() {
//        new NPETest().foo();
//    }
//    
//    private void foo() {
//        bar();
//    }
//    
//    private void bar() {
//        Object x = null;
//        synchronized(x) {
//            System.out.println("BAD!");
//        }
//    }
    
}
