public class StringBuilderTest {
    
    public static void main(String[] args) {
        //String s = "abc" + nextThreadNum();
        //System.out.println(s);
        
        StringBuilder sb = new StringBuilder();
        sb.append("abc");
        System.out.println(sb.toString());
    }
    
    private static int threadInitNumber;
    private static synchronized int nextThreadNum() {
        return threadInitNumber++;
    }
    
}
