package instructions;

public class TableSwitch {
    
    public static void main(String[] args) {
        switch (args.length) {
            case 3: System.out.println("3"); break;
            case 4: System.out.println("4"); break;
            case 5: System.out.println("5"); break;
            default: System.out.println("default");
        }
    }
    
}
