import {useParams} from '@remix-run/react'

export default function PersonalizedPortfolio() {
  const params = useParams()
  return <div>Personalized portfolio page with ID: {params.id} </div>
}
