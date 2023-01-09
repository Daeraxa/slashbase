import styles from './projectcard.module.scss'
import React, { useState } from 'react'
import { useAppDispatch } from '../../../redux/hooks'
import { createNewProject } from '../../../redux/projectsSlice'

type CreateNewProjectCardPropType = {}

const CreateNewProjectCard = (_: CreateNewProjectCardPropType) => {

    const [creating, setCreating] = useState(false)
    const [projectName, setProjectName] = useState('')
    const [loading, setLoading] = useState(false)

    const dispatch = useAppDispatch()

    const startCreatingProject = async () => {
        if (loading) {
            return
        }
        setLoading(true)
        await dispatch(createNewProject({ projectName }))
        setLoading(false)
        setCreating(false)
        setProjectName('')
    }

    return (
        <React.Fragment>
            {!creating &&
                <button
                    className="button"
                    onClick={() => {
                        setCreating(true)
                    }}>
                    <i className={"fas fa-folder-plus"} />
                    &nbsp;&nbsp;
                    Create New Project
                </button>
            }
            {
                creating &&
                <div className={"card " + styles.cardContainer}>
                    <div className="card-content">
                        <div className="field has-addons">
                            <div className="control is-expanded">
                                <input
                                    className="input"
                                    type="text"
                                    placeholder="Enter Project Name"
                                    value={projectName}
                                    onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setProjectName(e.target.value) }} />
                            </div>
                            <div className="control">
                                <button className="button is-primary" onClick={startCreatingProject}>
                                    {loading ? 'Creating' : 'Create'}
                                </button>
                                <button className={"delete " + styles.cancelBtn} onClick={() => { setCreating(false) }}></button>
                            </div>
                        </div>
                    </div>
                </div>
            }
        </React.Fragment>
    )
}


export default CreateNewProjectCard