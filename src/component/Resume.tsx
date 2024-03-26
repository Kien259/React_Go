import { useState } from "react";
import Title from "./Title";
import Education from "./Resume_Education";
import Skills from "./Resume_Skills";
import Experience from "./Resume_Experience";
import Achievement from "./Resume_Achievement";
import { FadeIn } from "./FadeIn";

const Resume = () => {
  const [skillData, setSkillData] = useState<Boolean>(true);
  const [educationData, setEducationData] = useState<Boolean>(false);
  const [experienceData, setExperienceData] = useState<Boolean>(false);
  const [achievementData, setAchievementData] = useState<Boolean>(false);
  return (
    <section
      id="resume"
      className="w-full py-20 border-b-[1px] border-b-gray-700"
    >
      <FadeIn>
        <div className="flex justify-center items-center text-center">
          <Title title="7+ YEARS OF EXPERIENCE" des="My Resume" />
        </div>
        <div>
          <ul className="w-full grid grid-cols-1 md:grid-cols-2 xl:grid-cols-4">
            <li
              onClick={() => {
                setSkillData(true);
                setEducationData(false);
                setExperienceData(false);
                setAchievementData(false);
              }}
              className={`${
                skillData
                  ? "border-designColor rounded-lg"
                  : "border-transparent"
              } resumeLi`}
            >
              Tech Stack
            </li>
            <li
              onClick={() => {
                setSkillData(false);
                setEducationData(true);
                setExperienceData(false);
                setAchievementData(false);
              }}
              className={`${
                educationData
                  ? "border-designColor rounded-lg"
                  : "border-transparent"
              } resumeLi`}
            >
              Education
            </li>
            <li
              onClick={() => {
                setSkillData(false);
                setEducationData(false);
                setExperienceData(true);
                setAchievementData(false);
              }}
              className={`${
                experienceData
                  ? "border-designColor rounded-lg"
                  : "border-transparent"
              } resumeLi`}
            >
              Experience
            </li>
            <li
              onClick={() => {
                setSkillData(false);
                setEducationData(false);
                setExperienceData(false);
                setAchievementData(true);
              }}
              className={`${
                achievementData
                  ? "border-designColor rounded-lg"
                  : "border-transparent"
              } resumeLi`}
            >
              Achievements
            </li>
          </ul>
        </div>
        {educationData && <Education />}
        {skillData && <Skills />}
        {achievementData && <Achievement />}
        {experienceData && <Experience />}
      </FadeIn>
    </section>
  );
};

export default Resume;