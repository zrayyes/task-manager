import React, { useEffect, useState } from "react";
import axios from "axios";
import './App.css';

interface Task {
  id: string;
  title: string;
  body: string;
  completed: boolean;
  updated_at: string;
  created_at: string;
}

const App: React.FC = () => {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [loading, setLoading] = useState(true);
  const [newTask, setNewTask] = useState({ title: "", body: "" });

  useEffect(() => {
    const fetchTasks = async () => {
      try {
        const response = await axios.get("http://localhost:8080/tasks");
        const sortedTasks = response.data.sort(
          (a: Task, b: Task) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime()
        );
        setTasks(sortedTasks);
      } catch (error) {
        console.error("Error fetching tasks:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchTasks();
  }, []);

  const toggleCompletion = async (task: Task) => {
    try {
      const updatedTask = { ...task, completed: !task.completed };
      await axios.put(`http://localhost:8080/tasks/${task.id}`, updatedTask);
      setTasks((prevTasks) =>
        prevTasks.map((t) => (t.id === task.id ? updatedTask : t))
      );
    } catch (error) {
      console.error("Error updating task:", error);
    }
  };

  const addTask = async () => {
    if (!newTask.title || !newTask.body) return;
    try {
      const response = await axios.post("http://localhost:8080/tasks", {
        ...newTask,
        completed: false,
      });
      setTasks((prevTasks) => [response.data, ...prevTasks]);
      setNewTask({ title: "", body: "" });
    } catch (error) {
      console.error("Error adding task:", error);
    }
  };

  if (loading) {
    return <div className="flex justify-center items-center h-screen">Loading...</div>;
  }

  return (
    <div className="min-h-screen bg-gray-100 p-4">
      <div className="max-w-2xl mx-auto bg-white shadow-md rounded-lg p-6">
        <h1 className="text-2xl font-bold mb-4 text-center text-gray-800">Task Manager</h1>
        <div className="mb-4">
          <input
            type="text"
            placeholder="Task Title"
            value={newTask.title}
            onChange={(e) => setNewTask({ ...newTask, title: e.target.value })}
            className="w-full p-2 mb-2 border border-gray-300 rounded bg-white text-gray-900"
          />
          <textarea
            placeholder="Task Body"
            value={newTask.body}
            onChange={(e) => setNewTask({ ...newTask, body: e.target.value })}
            className="w-full p-2 mb-2 border border-gray-300 rounded bg-white text-gray-900"
          />
          <button
            onClick={addTask}
            className="w-full bg-blue-500 text-white py-2 rounded hover:bg-blue-600"
          >
            Add Task
          </button>
        </div>
        {tasks.length === 0 ? (
          <p className="text-center text-gray-500">No tasks available.</p>
        ) : (
          <ul className="space-y-4">
            {tasks.map((task) => (
              <li
                key={task.id}
                className="flex items-center justify-between p-4 bg-gray-50 rounded-lg shadow-sm"
              >
                <div>
                  <h2 className="text-lg font-semibold text-gray-900">{task.title}</h2>
                  <p className="text-sm text-gray-600">{task.body}</p>
                </div>
                <div className="flex items-center space-x-2">
                  <input
                    type="checkbox"
                    checked={task.completed}
                    onChange={() => toggleCompletion(task)}
                    className="w-5 h-5 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                  />
                </div>
              </li>
            ))}
          </ul>
        )}
      </div>
    </div>
  );
};

export default App;
