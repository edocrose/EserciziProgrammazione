public abstract class Contenitore implements Comparable<Contenitore> {
    private double capienza, contenuto;
    private String liquido;

    public Contenitore(double contenuto, String liquido, double capienza)
            throws IllegalArgumentException, ExceededCapacityException {
        if (liquido == null || liquido.equals("")) {
            throw new IllegalArgumentException("Nome liquido nullo o vuoto");
        }
        if (contenuto <= 0) {
            throw new IllegalArgumentException("contenuto negativo");
        }
        if (capienza <= 0) {
            throw new IllegalArgumentException("contenuto negativo");
        }
        if (capienza < contenuto) {
            throw new ExceededCapacityException("capienza < contenuto");
        }
        this.contenuto = contenuto;
        this.capienza = capienza;
        this.liquido = liquido;
    }

    public double getCapienza() {
        return this.capienza;
    }

    public double getContenuto() {
        return this.contenuto;
    }

    public String getLiquido() {
        return this.liquido;
    }

    public void versa(Contenitore c) throws IncompatibleLiquidException {
        if (this.contenuto == 0) {
            return;
        }
        if (c.capienza == c.contenuto) {
            return;
        }

        if (c.liquido.equals("")) {
            c.liquido = this.liquido;
        }

        if (!(c.liquido.equals(this.liquido))) {
            throw new IncompatibleLiquidException("Liquidi non compatibili");
        }

        if (this.contenuto + c.contenuto < c.capienza) {
            c.contenuto += this.contenuto;
            this.contenuto = 0;
            this.liquido = "";
        } else {
            this.contenuto = (this.contenuto + c.contenuto) - c.capienza;
            c.contenuto = c.capienza;
        }
    }

    @Override
    public int compareTo(Contenitore o) {
        if (this.capienza > o.capienza) {
            return -1;
        }
        if (this.capienza < o.capienza) {
            return 1;
        }
        return 0;
    }

}
