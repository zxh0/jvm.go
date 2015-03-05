public class SysProps {
    
    public static void main(String[] args) {
        String[] keys = {
            "user.dir",
            "java.home",
            "java.class.path",
            "file.encoding",
            "sun.stdout.encoding",
            "sun.stderr.encoding",
            "file.separator",
            "path.separator",
            "line.separator",
        };
        
        for (String key : keys) {
            String val = System.getProperty(key);
            System.out.println(key + ":" + val);
        }
    }
    
}
