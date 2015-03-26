package helper;

import java.lang.reflect.Field;

public class ReflectionHelper {
    
    public static Object getStaticFieldValue(Class<?> c, String fieldName) throws ReflectiveOperationException {
        return getField(c, fieldName).get(null);
    }
    
    public static Object getFieldValue(Object obj, String fieldName) throws ReflectiveOperationException {
        return getField(obj.getClass(), fieldName).get(obj);
    }
    
    private static Field getField(Class<?> klass, String fieldName) {
        for (Class<?> c = klass; c != null; c = c.getSuperclass()) {
            try {
                Field f = c.getDeclaredField(fieldName);
                f.setAccessible(true);
                return f;
            } catch (NoSuchFieldException e) {
                // ignored
            }
        }
        return null;
    }
    
}
