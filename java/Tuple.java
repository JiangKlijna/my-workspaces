public class Tuple {
    private Tuple() {
    }

    public static class Zero implements java.io.Serializable {
        public static final int N = 0;

        private Zero() {
        }

        public Zero copy() {
            return new Zero();
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            return true;
        }

        @Override
        public int hashCode() {
            return Zero.class.hashCode();
        }

        @Override
        public String toString() {
            return "Tuple.Zero()";
        }
    }

    public static class One<A> extends Zero {
        public static final int N = 1;
        public final A a;

        private One(A a) {
            this.a = a;
        }

        public A getFirst() {
            return a;
        }

        public One<A> setFirst(A a) {
            return new One<A>(a);
        }

        @Override
        public One<A> copy() {
            return new One<A>(a);
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            One other = (One) o;
            if (!Tuple.equals(a, other.a)) return false;
            return true;
        }

        @Override
        public int hashCode() {
            int result = super.hashCode();
            result = 31 * result + (a != null ? a.hashCode() : 0);
            return result;
        }

        @Override
        public String toString() {
            return "Tuple.One(" + a + ")";
        }
    }

    public static class Two<A, B> extends One<A> {
        public static final int N = 2;
        public final B b;

        private Two(A a, B b) {
            super(a);
            this.b = b;
        }

        public B getSecond() {
            return b;
        }

        @Override
        public Two<A, B> setFirst(A a) {
            return new Two<A, B>(a, b);
        }

        public Two<A, B> setSecond(B b) {
            return new Two<A, B>(a, b);
        }

        @Override
        public Two<A, B> copy() {
            return new Two<A, B>(a, b);
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            Two other = (Two) o;
            if (!Tuple.equals(a, other.a)) return false;
            if (!Tuple.equals(b, other.b)) return false;
            return true;
        }

        @Override
        public int hashCode() {
            int result = super.hashCode();
            result = 31 * result + (b != null ? b.hashCode() : 0);
            return result;
        }

        @Override
        public String toString() {
            return "Tuple.Two(" + a + ", " + b + ")";
        }
    }

    public static class Three<A, B, C> extends Two<A, B> {
        public static final int N = 3;
        public final C c;

        private Three(A a, B b, C c) {
            super(a, b);
            this.c = c;
        }

        public C getThird() {
            return c;
        }

        @Override
        public Three<A, B, C> setFirst(A a) {
            return new Three<A, B, C>(a, b, c);
        }

        @Override
        public Three<A, B, C> setSecond(B b) {
            return new Three<A, B, C>(a, b, c);
        }

        public Three<A, B, C> setThird(C c) {
            return new Three<A, B, C>(a, b, c);
        }

        @Override
        public Three<A, B, C> copy() {
            return new Three<A, B, C>(a, b, c);
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            Three other = (Three) o;
            if (!Tuple.equals(a, other.a)) return false;
            if (!Tuple.equals(b, other.b)) return false;
            if (!Tuple.equals(c, other.c)) return false;
            return true;
        }

        @Override
        public int hashCode() {
            int result = super.hashCode();
            result = 31 * result + (c != null ? c.hashCode() : 0);
            return result;
        }

        @Override
        public String toString() {
            return "Tuple.Three(" + a + ", " + b + ", " + c + ")";
        }
    }

    public static class Four<A, B, C, D> extends Three<A, B, C> {
        public static final int N = 4;
        public final D d;

        private Four(A a, B b, C c, D d) {
            super(a, b, c);
            this.d = d;
        }

        public D getFourth() {
            return d;
        }

        @Override
        public Four<A, B, C, D> setFirst(A a) {
            return Tuple.make(a, b, c, d);
        }

        @Override
        public Four<A, B, C, D> setSecond(B b) {
            return Tuple.make(a, b, c, d);
        }

        @Override
        public Four<A, B, C, D> setThird(C c) {
            return Tuple.make(a, b, c, d);
        }

        public Four<A, B, C, D> setFourth(D d) {
            return Tuple.make(a, b, c, d);
        }

        @Override
        public Four<A, B, C, D> copy() {
            return Tuple.make(a, b, c, d);
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            Four other = (Four) o;
            if (!Tuple.equals(a, other.a)) return false;
            if (!Tuple.equals(b, other.b)) return false;
            if (!Tuple.equals(c, other.c)) return false;
            if (!Tuple.equals(d, other.d)) return false;
            return true;
        }

        @Override
        public int hashCode() {
            int result = super.hashCode();
            result = 31 * result + (d != null ? d.hashCode() : 0);
            return result;
        }

        @Override
        public String toString() {
            return "Tuple.Four(" + a + ", " + b + ", " + c + ", " + d + ")";
        }
    }

    public static class Five<A, B, C, D, E> extends Four<A, B, C, D> {
        public static final int N = 5;
        public final E e;

        private Five(A a, B b, C c, D d, E e) {
            super(a, b, c, d);
            this.e = e;
        }

        public E getFifth() {
            return e;
        }

        @Override
        public Five<A, B, C, D, E> setFirst(A a) {
            return new Five<A, B, C, D, E>(a, b, c, d, e);
        }

        @Override
        public Five<A, B, C, D, E> setSecond(B b) {
            return new Five<A, B, C, D, E>(a, b, c, d, e);
        }

        @Override
        public Five<A, B, C, D, E> setThird(C c) {
            return new Five<A, B, C, D, E>(a, b, c, d, e);
        }

        @Override
        public Five<A, B, C, D, E> setFourth(D d) {
            return new Five<A, B, C, D, E>(a, b, c, d, e);
        }

        public Five<A, B, C, D, E> setFifth(E e) {
            return new Five<A, B, C, D, E>(a, b, c, d, e);
        }

        @Override
        public Five<A, B, C, D, E> copy() {
            return new Five<A, B, C, D, E>(a, b, c, d, e);
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            Five other = (Five) o;
            if (!Tuple.equals(a, other.a)) return false;
            if (!Tuple.equals(b, other.b)) return false;
            if (!Tuple.equals(c, other.c)) return false;
            if (!Tuple.equals(d, other.d)) return false;
            if (!Tuple.equals(e, other.e)) return false;
            return true;
        }

        @Override
        public int hashCode() {
            int result = super.hashCode();
            result = 31 * result + (e != null ? e.hashCode() : 0);
            return result;
        }

        @Override
        public String toString() {
            return "Tuple.Five(" + a + ", " + b + ", " + c + ", " + d + ", " + e + ")";
        }
    }

    public static class Six<A, B, C, D, E, F> extends Five<A, B, C, D, E> {
        public static final int N = 6;
        public final F f;

        private Six(A a, B b, C c, D d, E e, F f) {
            super(a, b, c, d, e);
            this.f = f;
        }

        public F getSixth() {
            return f;
        }

        @Override
        public Six<A, B, C, D, E, F> setFirst(A a) {
            return new Six<A, B, C, D, E, F>(a, b, c, d, e, f);
        }

        @Override
        public Six<A, B, C, D, E, F> setSecond(B b) {
            return new Six<A, B, C, D, E, F>(a, b, c, d, e, f);
        }

        @Override
        public Six<A, B, C, D, E, F> setThird(C c) {
            return new Six<A, B, C, D, E, F>(a, b, c, d, e, f);
        }

        @Override
        public Six<A, B, C, D, E, F> setFourth(D d) {
            return new Six<A, B, C, D, E, F>(a, b, c, d, e, f);
        }

        @Override
        public Six<A, B, C, D, E, F> setFifth(E e) {
            return new Six<A, B, C, D, E, F>(a, b, c, d, e, f);
        }

        public Six<A, B, C, D, E, F> setSixth(F f) {
            return new Six<A, B, C, D, E, F>(a, b, c, d, e, f);
        }

        @Override
        public Six<A, B, C, D, E, F> copy() {
            return new Six<A, B, C, D, E, F>(a, b, c, d, e, f);
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            Six other = (Six) o;
            if (!Tuple.equals(a, other.a)) return false;
            if (!Tuple.equals(b, other.b)) return false;
            if (!Tuple.equals(c, other.c)) return false;
            if (!Tuple.equals(d, other.d)) return false;
            if (!Tuple.equals(e, other.e)) return false;
            if (!Tuple.equals(f, other.f)) return false;
            return true;
        }

        @Override
        public int hashCode() {
            int result = super.hashCode();
            result = 31 * result + (f != null ? f.hashCode() : 0);
            return result;
        }

        @Override
        public String toString() {
            return "Tuple.Six(" + a + ", " + b + ", " + c + ", " + d + ", " + e + ", " + f + ")";
        }
    }

    public static class Seven<A, B, C, D, E, F, G> extends Six<A, B, C, D, E, F> {
        public static final int N = 7;
        public final G g;

        private Seven(A a, B b, C c, D d, E e, F f, G g) {
            super(a, b, c, d, e, f);
            this.g = g;
        }

        public G getSeventh() {
            return g;
        }

        @Override
        public Seven<A, B, C, D, E, F, G> setFirst(A a) {
            return new Seven<A, B, C, D, E, F, G>(a, b, c, d, e, f, g);
        }

        @Override
        public Seven<A, B, C, D, E, F, G> setSecond(B b) {
            return new Seven<A, B, C, D, E, F, G>(a, b, c, d, e, f, g);
        }

        @Override
        public Seven<A, B, C, D, E, F, G> setThird(C c) {
            return new Seven<A, B, C, D, E, F, G>(a, b, c, d, e, f, g);
        }

        @Override
        public Seven<A, B, C, D, E, F, G> setFourth(D d) {
            return new Seven<A, B, C, D, E, F, G>(a, b, c, d, e, f, g);
        }

        @Override
        public Seven<A, B, C, D, E, F, G> setFifth(E e) {
            return new Seven<A, B, C, D, E, F, G>(a, b, c, d, e, f, g);
        }

        @Override
        public Seven<A, B, C, D, E, F, G> setSixth(F f) {
            return new Seven<A, B, C, D, E, F, G>(a, b, c, d, e, f, g);
        }

        public Seven<A, B, C, D, E, F, G> setSeventh(G g) {
            return new Seven<A, B, C, D, E, F, G>(a, b, c, d, e, f, g);
        }

        @Override
        public Seven<A, B, C, D, E, F, G> copy() {
            return new Seven<A, B, C, D, E, F, G>(a, b, c, d, e, f, g);
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            Seven other = (Seven) o;
            if (!Tuple.equals(a, other.a)) return false;
            if (!Tuple.equals(b, other.b)) return false;
            if (!Tuple.equals(c, other.c)) return false;
            if (!Tuple.equals(d, other.d)) return false;
            if (!Tuple.equals(e, other.e)) return false;
            if (!Tuple.equals(f, other.f)) return false;
            if (!Tuple.equals(g, other.g)) return false;
            return true;
        }

        @Override
        public int hashCode() {
            int result = super.hashCode();
            result = 31 * result + (g != null ? g.hashCode() : 0);
            return result;
        }

        @Override
        public String toString() {
            return "Tuple.Seven(" + a + ", " + b + ", " + c + ", " + d + ", " + e + ", " + f + ", " + g + ")";
        }
    }

    public static class Eight<A, B, C, D, E, F, G, H> extends Seven<A, B, C, D, E, F, G> {
        public static final int N = 8;
        public final H h;

        private Eight(A a, B b, C c, D d, E e, F f, G g, H h) {
            super(a, b, c, d, e, f, g);
            this.h = h;
        }

        public H getEighth() {
            return h;
        }

        @Override
        public Eight<A, B, C, D, E, F, G, H> setFirst(A a) {
            return new Eight<A, B, C, D, E, F, G, H>(a, b, c, d, e, f, g, h);
        }

        @Override
        public Eight<A, B, C, D, E, F, G, H> setSecond(B b) {
            return new Eight<A, B, C, D, E, F, G, H>(a, b, c, d, e, f, g, h);
        }

        @Override
        public Eight<A, B, C, D, E, F, G, H> setThird(C c) {
            return new Eight<A, B, C, D, E, F, G, H>(a, b, c, d, e, f, g, h);
        }

        @Override
        public Eight<A, B, C, D, E, F, G, H> setFourth(D d) {
            return new Eight<A, B, C, D, E, F, G, H>(a, b, c, d, e, f, g, h);
        }

        @Override
        public Eight<A, B, C, D, E, F, G, H> setFifth(E e) {
            return new Eight<A, B, C, D, E, F, G, H>(a, b, c, d, e, f, g, h);
        }

        @Override
        public Eight<A, B, C, D, E, F, G, H> setSixth(F f) {
            return new Eight<A, B, C, D, E, F, G, H>(a, b, c, d, e, f, g, h);
        }

        @Override
        public Eight<A, B, C, D, E, F, G, H> setSeventh(G g) {
            return new Eight<A, B, C, D, E, F, G, H>(a, b, c, d, e, f, g, h);
        }

        public Eight<A, B, C, D, E, F, G, H> setEighth(H h) {
            return new Eight<A, B, C, D, E, F, G, H>(a, b, c, d, e, f, g, h);
        }

        @Override
        public Eight<A, B, C, D, E, F, G, H> copy() {
            return new Eight<A, B, C, D, E, F, G, H>(a, b, c, d, e, f, g, h);
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            Eight other = (Eight) o;
            if (!Tuple.equals(a, other.a)) return false;
            if (!Tuple.equals(b, other.b)) return false;
            if (!Tuple.equals(c, other.c)) return false;
            if (!Tuple.equals(d, other.d)) return false;
            if (!Tuple.equals(e, other.e)) return false;
            if (!Tuple.equals(f, other.f)) return false;
            if (!Tuple.equals(g, other.g)) return false;
            if (!Tuple.equals(h, other.h)) return false;
            return true;
        }

        @Override
        public int hashCode() {
            int result = super.hashCode();
            result = 31 * result + (h != null ? h.hashCode() : 0);
            return result;
        }

        @Override
        public String toString() {
            return "Tuple.Eight(" + a + ", " + b + ", " + c + ", " + d + ", " + e + ", " + f + ", " + g + ", " + h + ")";
        }
    }

    public static class Nine<A, B, C, D, E, F, G, H, I> extends Eight<A, B, C, D, E, F, G, H> {
        public static final int N = 9;
        public final I i;

        private Nine(A a, B b, C c, D d, E e, F f, G g, H h, I i) {
            super(a, b, c, d, e, f, g, h);
            this.i = i;
        }

        public I getNinth() {
            return i;
        }

        @Override
        public Nine<A, B, C, D, E, F, G, H, I> setFirst(A a) {
            return new Nine<A, B, C, D, E, F, G, H, I>(a, b, c, d, e, f, g, h, i);
        }

        @Override
        public Nine<A, B, C, D, E, F, G, H, I> setSecond(B b) {
            return new Nine<A, B, C, D, E, F, G, H, I>(a, b, c, d, e, f, g, h, i);
        }

        @Override
        public Nine<A, B, C, D, E, F, G, H, I> setThird(C c) {
            return new Nine<A, B, C, D, E, F, G, H, I>(a, b, c, d, e, f, g, h, i);
        }

        @Override
        public Nine<A, B, C, D, E, F, G, H, I> setFourth(D d) {
            return new Nine<A, B, C, D, E, F, G, H, I>(a, b, c, d, e, f, g, h, i);
        }

        @Override
        public Nine<A, B, C, D, E, F, G, H, I> setFifth(E e) {
            return new Nine<A, B, C, D, E, F, G, H, I>(a, b, c, d, e, f, g, h, i);
        }

        @Override
        public Nine<A, B, C, D, E, F, G, H, I> setSixth(F f) {
            return new Nine<A, B, C, D, E, F, G, H, I>(a, b, c, d, e, f, g, h, i);
        }

        @Override
        public Nine<A, B, C, D, E, F, G, H, I> setSeventh(G g) {
            return new Nine<A, B, C, D, E, F, G, H, I>(a, b, c, d, e, f, g, h, i);
        }

        @Override
        public Nine<A, B, C, D, E, F, G, H, I> setEighth(H h) {
            return new Nine<A, B, C, D, E, F, G, H, I>(a, b, c, d, e, f, g, h, i);
        }

        public Nine<A, B, C, D, E, F, G, H, I> setNinth(I i) {
            return new Nine<A, B, C, D, E, F, G, H, I>(a, b, c, d, e, f, g, h, i);
        }

        @Override
        public Nine<A, B, C, D, E, F, G, H, I> copy() {
            return new Nine<A, B, C, D, E, F, G, H, I>(a, b, c, d, e, f, g, h, i);
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            Nine other = (Nine) o;
            if (!Tuple.equals(a, other.a)) return false;
            if (!Tuple.equals(b, other.b)) return false;
            if (!Tuple.equals(c, other.c)) return false;
            if (!Tuple.equals(d, other.d)) return false;
            if (!Tuple.equals(e, other.e)) return false;
            if (!Tuple.equals(f, other.f)) return false;
            if (!Tuple.equals(g, other.g)) return false;
            if (!Tuple.equals(h, other.h)) return false;
            if (!Tuple.equals(i, other.i)) return false;
            return true;
        }

        @Override
        public int hashCode() {
            int result = super.hashCode();
            result = 31 * result + (i != null ? i.hashCode() : 0);
            return result;
        }

        @Override
        public String toString() {
            return "Tuple.Nine(" + a + ", " + b + ", " + c + ", " + d + ", " + e + ", " + f + ", " + g + ", " + h + ", " + i + ")";
        }
    }

    public static class Ten<A, B, C, D, E, F, G, H, I, J> extends Nine<A, B, C, D, E, F, G, H, I> {
        public static final int N = 10;
        public final J j;

        private Ten(A a, B b, C c, D d, E e, F f, G g, H h, I i, J j) {
            super(a, b, c, d, e, f, g, h, i);
            this.j = j;
        }

        public J getTenth() {
            return j;
        }

        @Override
        public Ten<A, B, C, D, E, F, G, H, I, J> setFirst(A a) {
            return new Ten<A, B, C, D, E, F, G, H, I, J>(a, b, c, d, e, f, g, h, i, j);
        }

        @Override
        public Ten<A, B, C, D, E, F, G, H, I, J> setSecond(B b) {
            return new Ten<A, B, C, D, E, F, G, H, I, J>(a, b, c, d, e, f, g, h, i, j);
        }

        @Override
        public Ten<A, B, C, D, E, F, G, H, I, J> setThird(C c) {
            return new Ten<A, B, C, D, E, F, G, H, I, J>(a, b, c, d, e, f, g, h, i, j);
        }

        @Override
        public Ten<A, B, C, D, E, F, G, H, I, J> setFourth(D d) {
            return new Ten<A, B, C, D, E, F, G, H, I, J>(a, b, c, d, e, f, g, h, i, j);
        }

        @Override
        public Ten<A, B, C, D, E, F, G, H, I, J> setFifth(E e) {
            return new Ten<A, B, C, D, E, F, G, H, I, J>(a, b, c, d, e, f, g, h, i, j);
        }

        @Override
        public Ten<A, B, C, D, E, F, G, H, I, J> setSixth(F f) {
            return new Ten<A, B, C, D, E, F, G, H, I, J>(a, b, c, d, e, f, g, h, i, j);
        }

        @Override
        public Ten<A, B, C, D, E, F, G, H, I, J> setSeventh(G g) {
            return new Ten<A, B, C, D, E, F, G, H, I, J>(a, b, c, d, e, f, g, h, i, j);
        }

        @Override
        public Ten<A, B, C, D, E, F, G, H, I, J> setEighth(H h) {
            return new Ten<A, B, C, D, E, F, G, H, I, J>(a, b, c, d, e, f, g, h, i, j);
        }

        @Override
        public Ten<A, B, C, D, E, F, G, H, I, J> setNinth(I i) {
            return new Ten<A, B, C, D, E, F, G, H, I, J>(a, b, c, d, e, f, g, h, i, j);
        }

        public Ten<A, B, C, D, E, F, G, H, I, J> setTenth(J j) {
            return new Ten<A, B, C, D, E, F, G, H, I, J>(a, b, c, d, e, f, g, h, i, j);
        }

        @Override
        public Ten<A, B, C, D, E, F, G, H, I, J> copy() {
            return new Ten<A, B, C, D, E, F, G, H, I, J>(a, b, c, d, e, f, g, h, i, j);
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            Ten other = (Ten) o;
            if (!Tuple.equals(a, other.a)) return false;
            if (!Tuple.equals(b, other.b)) return false;
            if (!Tuple.equals(c, other.c)) return false;
            if (!Tuple.equals(d, other.d)) return false;
            if (!Tuple.equals(e, other.e)) return false;
            if (!Tuple.equals(f, other.f)) return false;
            if (!Tuple.equals(g, other.g)) return false;
            if (!Tuple.equals(h, other.h)) return false;
            if (!Tuple.equals(i, other.i)) return false;
            if (!Tuple.equals(j, other.j)) return false;
            return true;
        }

        @Override
        public int hashCode() {
            int result = super.hashCode();
            result = 31 * result + (j != null ? j.hashCode() : 0);
            return result;
        }

        @Override
        public String toString() {
            return "Tuple.Ten(" + a + ", " + b + ", " + c + ", " + d + ", " + e + ", " + f + ", " + g + ", " + h + ", " + i + ", " + j + ")";
        }
    }

    public static final Zero make() {
        return new Zero();
    }

    public static final <A> One<A> make(A a) {
        return new One<A>(a);
    }

    public static final <A, B> Two<A, B> make(A a, B b) {
        return new Two<A, B>(a, b);
    }

    public static final <A, B, C> Three<A, B, C> make(A a, B b, C c) {
        return new Three<A, B, C>(a, b, c);
    }

    public static final <A, B, C, D> Four<A, B, C, D> make(A a, B b, C c, D d) {
        return new Four<A, B, C, D>(a, b, c, d);
    }

    public static final <A, B, C, D, E> Five<A, B, C, D, E> make(A a, B b, C c, D d, E e) {
        return new Five<A, B, C, D, E>(a, b, c, d, e);
    }

    public static final <A, B, C, D, E, F> Six<A, B, C, D, E, F> make(A a, B b, C c, D d, E e, F f) {
        return new Six<A, B, C, D, E, F>(a, b, c, d, e, f);
    }

    public static final <A, B, C, D, E, F, G> Seven<A, B, C, D, E, F, G> make(A a, B b, C c, D d, E e, F f, G g) {
        return new Seven<A, B, C, D, E, F, G>(a, b, c, d, e, f, g);
    }

    public static final <A, B, C, D, E, F, G, H> Eight<A, B, C, D, E, F, G, H> make(A a, B b, C c, D d, E e, F f, G g, H h) {
        return new Eight<A, B, C, D, E, F, G, H>(a, b, c, d, e, f, g, h);
    }

    public static final <A, B, C, D, E, F, G, H, I> Nine<A, B, C, D, E, F, G, H, I> make(A a, B b, C c, D d, E e, F f, G g, H h, I i) {
        return new Nine<A, B, C, D, E, F, G, H, I>(a, b, c, d, e, f, g, h, i);
    }

    public static final <A, B, C, D, E, F, G, H, I, J> Ten<A, B, C, D, E, F, G, H, I, J> make(A a, B b, C c, D d, E e, F f, G g, H h, I i, J j) {
        return new Ten<A, B, C, D, E, F, G, H, I, J>(a, b, c, d, e, f, g, h, i, j);
    }

    private static final boolean equals(Object a, Object b) {
        return (a == b) || (a != null && a.equals(b));
    }
}
