public class Cuboide extends Contenitore {
    private double a, b, c;

    public Cuboide(double contenuto, String liquido, double capienza, double a, double b, double c)
            throws IllegalArgumentException, ExceededCapacityException {
        super(contenuto, liquido, capienza);
        if (a <= 0) {
            throw new IllegalArgumentException("a negativo");
        }
        if (b <= 0) {
            throw new IllegalArgumentException("b negativo");
        }
        if (c <= 0) {
            throw new IllegalArgumentException("c negativo");
        }
        this.a = a;
        this.b = b;
        this.c = c;
    }

}
