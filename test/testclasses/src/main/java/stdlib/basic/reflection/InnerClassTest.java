package stdlib.basic.reflection;

public class InnerClassTest {

    private static class Inner {}

    public static void main(String[] args) {
        System.out.println(InnerClassTest.class.getModifiers()); // 1
        System.out.println(Inner.class.getModifiers()); // 10
    }

}
