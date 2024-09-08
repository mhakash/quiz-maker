import { json, LoaderFunctionArgs } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";

export type Tag = {
  id: number;
  name: string;
};

export type Quiz = {
  id: number;
  question: string;
  options: {
    [key: string]: string;
  };
  answer: string;
  tags: Tag[] | null;
};

const getQuizzes = async () => {
  const response = await fetch("http://127.0.0.1:1234/quiz");
  const data: { quizzes: Quiz[] } = await response.json();
  return data;
};

export async function loader({ request }: LoaderFunctionArgs) {
  const quizzes = await getQuizzes();
  return json(quizzes);
}

export default function QuizzesPage() {
  const data = useLoaderData<typeof loader>();

  return (
    <>
      <div className="my-8">
        {data.quizzes.map((quiz) => (
          <Quiz key={quiz.id} quiz={quiz} />
        ))}
      </div>
    </>
  );
}

const Quiz = (props: { quiz: Quiz }) => {
  return (
    <div className="py-1 px-3 my-1 border border-gray-200 rounded">
      <div className="mb-1">
        {props.quiz.question}
        {props.quiz.tags?.map((tag) => (
          <Tag key={tag.id} tag={tag} />
        ))}
      </div>
      <ul>
        {Object.entries(props.quiz.options).map(([key, value]) => (
          <li key={key}>
            <label>
              <input
                type="radio"
                name={props.quiz.question}
                value={key}
                className="mr-2"
              />
              {value}
            </label>
          </li>
        ))}
      </ul>
    </div>
  );
};

const Tag = (props: { tag: Tag }) => {
  return (
    <span className="bg-green-300 text-gray-800 text-xs px-2 rounded-md mx-1">
      {props.tag.name}
    </span>
  );
};
