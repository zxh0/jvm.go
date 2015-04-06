package java6.instructions;

public class ANewArray {
    
    public static void main(String[] args) {
        test1();
        test2();
    }
    
    private static void test1() {
        Object[] arr = new Object[8];
        if (arr.length == 8) {
            System.out.println("OK!");
        } else {
            System.out.println("Fail!");
        }
    }
    
    private static void test2() {
        int[][][] y = {
            {
                {1},
                {1, 2},
                {1, 2, 3}
            }
        };
        System.out.println(y.length);
        System.out.println(y[0].length);
        System.out.println(y[0][0].length);
        System.out.println(y[0][1].length);
        System.out.println(y[0][2].length);
    }
    
}
