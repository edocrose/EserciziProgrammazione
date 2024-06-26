public class Decorazione {
    private String nome;
    private double peso;

    public Decorazione(String nome, double peso) throws IllegalArgumentException {
        if (nome == null || nome.equals("")) {
            throw new IllegalArgumentException("Nome nullo");
        }
        if (peso <= 0) {
            throw new IllegalArgumentException("Peso <= 0");
        }
        this.nome = nome;
        this.peso = peso;
    }

    public String getNome() {
        return this.nome;
    }

    public double getPeso() {
        return this.peso;
    }
}
