import { useEffect, useState } from "react";
import { useParams } from "react-router";
import { getIpInfo, redirection } from "../lib/fetch";

export default function Redirect() {
  const { short_code } = useParams();
  const [ipInfo, setIpInfo] = useState(null);
  const [runOnce, setRunOnce] = useState(false);

  useEffect(() => {
    getIpInfo().then((data) => {
      setIpInfo(data);
    });
  }, []);

  useEffect(() => {
    if (ipInfo) {
      const request = {
        short_code: short_code,
        ip_address: ipInfo?.ip,
        location: ipInfo?.loc,
        city: ipInfo?.city,
      };
      if (!runOnce) {
        redirection(request)
          .then((data) => {
            window.location.replace(data.origin_url);
          })
          .catch((err) => {
            console.log(err);
          });
        setRunOnce(true);
      }
    }
  }, [ipInfo, short_code]);

  return <div>Redirecting...</div>;
}
