import {
  http
} from "../utils/http";

const MonitoringService = {
  async get() {
    var result = await http.get("/api/post/getposts");
    if (result.status === 200) {
      return result.data;
    } else {
      console.error(result.error);
      throw result.error;
    }
  },
  async save(value) {
    var result = await http.post("/api/post/newpost", value);
    if (result.status === 200) {
      return result.data;
    } else {
      console.error(result.error);
      throw result.error;
    }
  }
};

export default MonitoringService;
