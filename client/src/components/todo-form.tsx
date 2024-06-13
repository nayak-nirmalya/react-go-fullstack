import { useState } from "react";
import { Button, Flex, Input, Spinner } from "@chakra-ui/react";
import { IoMdAdd } from "react-icons/io";

export function ToDoForm() {
  const [newTodo, setNewTodo] = useState("");
  const [isPending, setIsPending] = useState(false);

  const createTodo = async (e: React.FormEvent) => {
    e.preventDefault();
    alert("ToDo Added!");
  };
  return (
    <form onSubmit={createTodo}>
      <Flex gap={2}>
        <Input
          type="text"
          value={newTodo}
          onChange={(e) => setNewTodo(e.target.value)}
          ref={(input) => input && input.focus()}
        />
        <Button
          mx={2}
          type="submit"
          _active={{
            transform: "scale(.97)",
          }}
        >
          {isPending ? <Spinner size={"xs"} /> : <IoMdAdd size={30} />}
        </Button>
      </Flex>
    </form>
  );
}
