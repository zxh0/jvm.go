public class SysProps {
    
    public static void main(String[] args) {
        String[] keys = {
            "file.encoding",
            "sun.stdout.encoding",
            "sun.stderr.encoding",
            "file.separator",
            "path.separator",
            "line.separator",
            "java.home",
        };
        
        for (String key : keys) {
            String val = System.getProperty(key);
            System.out.println(key + ":" + val);
        }
    }
    
}
