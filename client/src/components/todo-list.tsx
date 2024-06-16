import { Flex, Spinner, Stack, Text } from "@chakra-ui/react";
import { useQuery } from "@tanstack/react-query";

import { ToDoItem } from "./todo-item";

import { ToDo } from "../types";

export function ToDoList() {
  const { data: todos, isLoading } = useQuery<ToDo[]>({
    queryKey: ["todos"],
    queryFn: async () => {
      try {
        const res = await fetch(BASE_URL + "/todos");
        const data = await res.json();

        if (!res.ok) {
          throw new Error(data.error || "Something Went Wrong! ❌");
        }

        return data || [];
      } catch (error) {
        console.error(error);
      }
    },
  });

  return (
    <>
      <Text
        fontSize={"4xl"}
        textTransform={"uppercase"}
        fontWeight={"bold"}
        textAlign={"center"}
        my={2}
        bgGradient="linear(to-l, #0b85f8, #00ffff)"
        bgClip="text"
      >
        {`Today's Task${todos?.length === 1 ? "" : "s"}`}
      </Text>
      {isLoading && (
        <Flex justifyContent={"center"} my={4}>
          <Spinner size={"xl"} />
        </Flex>
      )}
      {!isLoading && todos?.length === 0 && (
        <Stack alignItems={"center"} gap="3">
          <Text fontSize={"xl"} textAlign={"center"} color={"gray.500"}>
            All Tasks Completed! 🎉
          </Text>
          <img src="/go.png" alt="Go logo" width={70} height={70} />
        </Stack>
      )}
      <Stack gap={3}>
        {todos?.map((todo) => (
          <ToDoItem key={todo.id} todo={todo} />
        ))}
      </Stack>
    </>
  );
}
