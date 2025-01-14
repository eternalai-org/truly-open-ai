import "./App.css";
import { Studio } from "@agent-studio/studio-dnd";
import getAgentModelCategories from "./constants/categories";

const args = {
  categories: getAgentModelCategories("create"),
  dataSource: {},
  showConnectLine: true,
};

function App() {
  return (
    <div className="studio-app">
      {/* @ts-ignore */}
      <Studio {...args} data={[]} onChange={(data) => {}} />
    </div>
  );
}

export default App;
