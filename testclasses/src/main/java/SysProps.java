public class SysProps {
    
    public static void main(String[] args) {
        String val = System.getProperty("key", "val");
        System.out.println(val);
    }
    
}
