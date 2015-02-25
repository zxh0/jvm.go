public class RuntimeTest {
    
    public static void main(String[] args) {
        Runtime rt = Runtime.getRuntime();
        System.out.println("availableProcessors:" + rt.availableProcessors());
        System.out.println("freeMemory:" + rt.freeMemory());
    }
    
}
