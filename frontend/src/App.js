import React from "react";
import { Router, Switch, Route } from "react-router-dom";
import { history } from "./history";
import Sum from "./routes/Sum";

function App() {
  return (
    <Router history={history}>
      <Switch>
        <Route exact path="/" component={Sum}></Route>
      </Switch>
    </Router>
  );
}

export default App;
