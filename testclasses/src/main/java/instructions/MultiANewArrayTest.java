package instructions;

public class MultiANewArrayTest {
    
    public static void main(String[] args) {
        int[][][] x = new int[2][3][5];
        System.out.println(x.length);
        System.out.println(x[0].length);
        System.out.println(x[1][2].length);
        
        x[1][2][3] = 7;
        System.out.println(x[1][2][3]);
    }
    
}
