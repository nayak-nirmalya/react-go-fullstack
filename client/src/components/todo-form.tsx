import { useState } from "react";
import { Button, Flex, Input, Spinner } from "@chakra-ui/react";
import { IoMdAdd } from "react-icons/io";
import { useMutation, useQueryClient } from "@tanstack/react-query";

import { BASE_URL } from "../const";

export function ToDoForm() {
  const queryClient = useQueryClient();

  const [newToDo, setNewToDo] = useState("");

  const { mutate: createToDo, isPending: isCreating } = useMutation({
    mutationKey: ["createTodo"],
    mutationFn: async (e: React.FormEvent) => {
      e.preventDefault();
      try {
        const res = await fetch(BASE_URL + `/todos`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ body: newToDo }),
        });
        const data = await res.json();

        if (!res.ok) {
          throw new Error(data.error || "Something Went Wrong! ❌");
        }

        setNewToDo("");

        return data;
      } catch (error) {
        console.error(error);
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
    onError: (error) => {
      // TODO: Show Error Toast Notification
      console.error(error);
    },
  });

  return (
    <form onSubmit={createToDo}>
      <Flex gap={2}>
        <Input
          type="text"
          value={newToDo}
          onChange={(e) => setNewToDo(e.target.value)}
          ref={(input) => input && input.focus()}
        />
        <Button
          mx={2}
          type="submit"
          _active={{
            transform: "scale(.97)",
          }}
        >
          {isCreating ? <Spinner size={"xs"} /> : <IoMdAdd size={30} />}
        </Button>
      </Flex>
    </form>
  );
}
