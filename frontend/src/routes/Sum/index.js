import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";

const Sum = () => {
  function useQuery() {
    return new URLSearchParams(useLocation().search);
  }

  const query = useQuery();

  const a = query.get("a");
  const b = query.get("b");
  const [error, setError] = useState(null);
  const [isLoaded, setIsLoaded] = useState(false);
  const [result, setResult] = useState(0);

  useEffect(() => {
    const body = JSON.stringify({ a: a, b: b });
    fetch("http://localhost:8000/api/sum", {
      method: "POST",
      body: body,
    })
      .then((res) => res.json())
      .then((result) => {
        setIsLoaded(true);
        setResult(result.hasil);
      })
      .catch((error) => {
        setIsLoaded(true);
        setError(error);
      });
  }, [a, b]);

  if (error) {
    return <div>Error: {error.message}</div>;
  } else if (!isLoaded) {
    return <div>Loading...</div>;
  } else {
    return <div>Result : {result}</div>;
  }
};

export default Sum;
