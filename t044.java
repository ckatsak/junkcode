/**
 * Figuring out how wait()/notify() mechanism works.
 * Implementing a schema where Christoulas and Mitsoulas and Tux take turns in a game.
 * The schema works as is, even after adding a new player in main().
 */
package t044;

/**
 * @author christos
 *
 */
public class t044 {

	/**
	 * @param args
	 */
	public static void main(String[] args) {
		Battlefield battlefield = new Battlefield();
		(new Thread(battlefield)).start();
		(new Thread(new Player("Christoulas", 0, battlefield))).start();
		(new Thread(new Player("Mitsoulas", 1, battlefield))).start();
		(new Thread(new Player("Tux", 2, battlefield))).start();
	}

}

class Player implements Runnable {
	
	private String name;
	private int playerID;
	private Battlefield battlefield;
	
	private static int numberOfPlayers = 0;
	
	public Player(String name, int playerID, Battlefield battlefield) {
		this.name = name;
		this.playerID = playerID;
		this.battlefield = battlefield;
		++Player.numberOfPlayers;
	}
	
	public String getName() {
		return this.name;
	}

	@Override
	public void run() {
		/* Loop to make sure the two threads will never stop running. */
		while (true) {
			/* Do all this atomically. */
			synchronized(this.battlefield) {
				/* Wait on the monitor object as long as it's not your turn.
				 * You'll get notified when the turn changes. */
				while (battlefield.getPlayerTurn() != this.playerID) {
					try {
						battlefield.wait();
					} catch (InterruptedException e1) {
						e1.printStackTrace();
					}
				}
				/* Now that it's your turn, do your job. */
				System.out.println(this.getName());
				/* Change the turn and notify . */
				battlefield.setPlayerTurn((this.playerID + 1) % Player.numberOfPlayers);
				/* Notify whoever waits for his turn. */
				battlefield.notifyAll();
				
				/* Rest for a sec. */
				try {
					Thread.sleep(1000);
				} catch (InterruptedException e) {
					e.printStackTrace();
				}
			}
		}
	}
	
}

class Battlefield implements Runnable {

	private int playerTurn;
	
	public Battlefield() {
		this.playerTurn = 0;
	}
	
	public int getPlayerTurn() {
		return this.playerTurn;
	}
	
	public void setPlayerTurn(int playerTurn) {
		this.playerTurn = playerTurn;
	}
	
	@Override
	public void run() {
		// TODO Auto-generated method stub
		
	}
	
}
