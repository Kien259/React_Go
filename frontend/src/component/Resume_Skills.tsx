import { logos } from "~/data/logos"

const Skills = () => {
  return (
    <div className="flex justify-around flex-wrap gap-8 p-16 bg-bodyColor">
      {logos.map(({_id, title, link}) => (
        <div key={_id} className="flex min-w-[250px] max-w-[250px] h-auto bg-black bg-opacity-25 text-gray-200 text-xl items-center justify-center rounded-md shadow-shadowOne hover:bg-opacity-40 hover:-translate-y-1 transition-all hover:text-designColor cursor-pointer duration-300 p-4 gap-4">
          <img src={link} alt={title} className="w-20 h-20" />
          <span>{title}</span>
        </div>
      ))}
    </div>
  );
};

export default Skills;

