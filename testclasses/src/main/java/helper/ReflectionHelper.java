package helper;

import java.lang.reflect.Field;

public class ReflectionHelper {
    
    public static Object getFieldValue(Object obj, String fieldName) throws ReflectiveOperationException {
        return getField(obj, fieldName).get(obj);
    }
    
    private static Field getField(Object obj, String fieldName) {
        for (Class<?> c = obj.getClass(); c != null; c = c.getSuperclass()) {
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
