import TopvarDesktop from "./TopvarDesktop"
import TopvarMobile from "./TopvarMobile"

type Props = {}

const Topvar = (props: Props) => {
  return (
    <div className="w-full">
      <div className="w-full hidden md:block">
        <TopvarDesktop />
      </div>
      <div className="w-full md:hidden">
        <TopvarMobile />
      </div>
    </div>
  )
}

export default Topvar