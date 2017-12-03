import java.util.Date;
import java.util.Calendar;

public class Times {

    public static void main(String[] args) {
        System.out.println(operate("+10y"));
        System.out.println(operate("-10M"));
        System.out.println(operate("+10d"));
        System.out.println(operate("-10H"));
        System.out.println(operate("+10m"));
        System.out.println(operate("-10s"));
    }

    private static final Unit[] units = Unit.class.getEnumConstants();

    private static final Unit match(char c) {
        for (Unit u : units) if (u.mChar == c) return u;
        return null;
    }

    private static final Unit match(int i) {
        for (Unit u : units) if (u.mUnit == i) return u;
        return null;
    }

    private enum Unit {
        y('y', Calendar.YEAR),
        M('M', Calendar.MONTH),
        d('d', Calendar.DATE),
        H('H', Calendar.HOUR),
        m('m', Calendar.MINUTE),
        s('s', Calendar.SECOND);

        private final char mChar;
        private final int mUnit;

        private Unit(char c, int i) {
            mChar = c;
            mUnit = i;
        }
    }

    public static final Date operate(String mode) {
        return operate(mode, Calendar.getInstance());
    }

    public static final Date operate(String mode, long millis) {
        Calendar r = Calendar.getInstance();
        r.setTimeInMillis(millis);
        return operate(mode, r);
    }

    public static final Date operate(String mode, Date date) {
        Calendar r = Calendar.getInstance();
        r.setTime(date);
        return operate(mode, r);
    }

    public static final Date operate(String mode, Calendar r) {
        Unit u = match(mode.charAt(mode.length() - 1));
        if (u == null) return r.getTime();
        int num;
        try {
            num = Integer.parseInt(mode.substring(0, mode.length() - 1));
        } catch (NumberFormatException e) {
            return r.getTime();
        }
        r.add(u.mUnit, num);
        return r.getTime();
    }
}
