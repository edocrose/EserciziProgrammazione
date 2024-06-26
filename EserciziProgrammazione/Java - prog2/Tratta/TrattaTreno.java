public class TrattaTreno extends Tratta {

    private double co2;

    public TrattaTreno(String origine, String destinazione, double lung, double vel, double co2)
            throws IllegalArgumentException {
        // MODIFIES: this
        // EFFECTS: inizializza una tratta in treno con gli attributi di Tratta + un
        // altro
        // se il nuovo parametro non è corretto lancia IllegalArgumentException
        super(origine, destinazione, lung, vel);

        if (co2 <= 0) {
            throw new IllegalArgumentException("quantità di CO2 <= 0");
        }

        this.co2 = co2;
    }

    public double getCo2() {
        return this.co2;
    }

    @Override
    public double calcolaCO2() {
        // EFFECTS: restituisce la quantità di CO2 prodotta come lung * co2
        return this.getLung() * this.co2;
    }

    @Override
    public String toString() {
        return "(Treno) Tratta da " + this.getOrigine() + " a " + this.getDestinazione() + "; lunghezza: "
                + this.getLung() + ", velocità: " + this.getVel() + ", CO2/km: " + this.co2;
    }

}
