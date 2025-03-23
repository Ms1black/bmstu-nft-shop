import { useState, useEffect } from "react";
import Web3 from "web3";
import "./index.css";

const App = () => {
  const [balance, setBalance] = useState("0");

  useEffect(() => {
    const loadBlockchainData = async () => {
      const web3 = new Web3("http://127.0.0.1:8545");
      const accounts = await web3.eth.getAccounts();
      if (accounts.length > 0) {
        const balanceWei = await web3.eth.getBalance(accounts[0]);
        const balanceEth = web3.utils.fromWei(balanceWei, "ether");
        setBalance(balanceEth);
      }
    };

    loadBlockchainData();
  }, []);

  return (
    <div className="container">
      <h2>Баланс пользователя:</h2>
      <p>{balance} ETH</p>
    </div>
  );
};

export default App;
