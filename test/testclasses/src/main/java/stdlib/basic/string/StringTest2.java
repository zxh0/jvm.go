package stdlib.basic.string;

public class StringTest2 {

    public static void main(String[] args) {
        String s = "\uD800";
        System.out.println(s.length());
        System.out.println((int)s.charAt(0)); // 55296
    }

}
