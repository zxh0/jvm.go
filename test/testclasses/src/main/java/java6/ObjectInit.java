package java6;


public class ObjectInit {

    char x;
    static long a;

    public static void main(String[] args) {
        ObjectInit o = new ObjectInit();
        char y = o.x;
        if (y == 0) {
            System.out.println("OK1!");
        } else {
            System.out.println("Fail1!");
        }

        long b = a;
        if (b == 0) {
            System.out.println("OK2!");
        } else {
            System.out.println("Fail2!");
        }
    }

}
