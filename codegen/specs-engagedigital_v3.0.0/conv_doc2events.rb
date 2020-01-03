#!ruby

require 'multi_json'

doc = %Q{api_access_token.created API access token created
api_access_token.destroyed API access token destroyed
api_access_token.updated API access token updated
automatic_exports_task.created Automatic export task created
automatic_exports_task.destroyed Automatic export task destroyed
automatic_exports_task.failed Automatic export task failed
automatic_exports_task.succeed Automatic export task succeed
automatic_exports_task.updated Automatic export task updated
category.created Category created
category.destroyed Category deleted
category.updated Category updated
community.created Community created
community.destroyed Community deleted
community.updated Community updated
content.admin_stamped Answer admin stamped
content.approved Message approved
content.author_blocked Identity blocked
content.author_stamped Answer author stamped
content.author_unblocked Identity unblocked
content.auto_categorization_infirmed ICE message categorization corrected
content.auto_categorization_not_precise ICE message categorization failed
       content.auto_categorized Message categorized by ICE
      content.auto_ignored Message auto archived
      content.categorized Message categorized
      content.destroyed Message deleted
      content.discussion_initiated Discussion initiated
      content.discussion_planned Discussion planned
      content.ice_nth_content_ignored ICE ignored a message
      content.ignored Message ignored
      content.imported Message importated
      content.liked Message liked
      content.moderated_ban Identity blocked
      content.moderated_modif Message edited
      content.moderated_nok Message unpublished
      content.moderated_ok Message approved
      content.planned_discussion_synchronized Planned discussion synchronized
      content.published Message published
      content.recategorized Message recategorized
      content.replied Message replied
      content.reply_assistant_used Reply assistant used to reply to a message
      content.retried_synchronization Message synchronization retried
      content.retweeted Message retweeted
      content.source_changed Source changed
      content.stared Message starred
      content.thread_closed Thread closed
      content.thread_opened Thread opened
      content.unliked Message unliked
     content.unpublished Message unpublished
       content.unstamped Message unstamped
      content.unstared Message unstarred
      content.updated Message updated
      content_source.created Content source created
      content_source.destroyed Content source destroyed
      content_source.updated Content source updated
      content_thread.categorized Thread categorized
      content_thread.destroyed Thread destroyed
      content_thread.recategorized Thread recategorized
      custom_field.created Custom field created
      custom_field.destroyed Custom field deleted
      custom_field.updated Custom field updated
      expired_data_purge.deleted Expired data deleted
      extension.created Extension added
      extension.destroyed Extension removed
      extension.updated Extension updated
      folder.created Folder created
      folder.destroyed Folder destroyed
      folder.updated Folder updated
      identity.followed Identity followed
      identity.unfollowed Identity unfollowed
      identity.updated Identity updated
      intervention.assigned Message assigned
      intervention.canceled Intervention canceled
      intervention.closed Message solved
      intervention.deferred Intervention deferred
     intervention.opened Message engaged
       intervention.recategorized Intervention recategorized
      intervention.reopened Intervention reopened
      intervention.updated Intervention updated
      intervention.user_updated Intervention's user updated
      intervention_comment.created Intervention commented
      intervention_comment.destroyed Intervention destroyed
     reply_assistant_knowledge_base_entry.created Reply assistant knowledge base entry created
     reply_assistant_knowledge_base_entry.destroyed Reply assistant knowledge base entry destroyed
      reply_assistant_knowledge_base_entry.updated Reply assistant knowledge base entry updated
     reply_assistant_knowledge_base_version.created Reply assistant knowledge base version created
     reply_assistant_knowledge_base_version.destroyed Reply assistant knowledge base version destroyed
      reply_assistant_knowledge_base_version.updated Reply assistant knowledge base version updated
      reply_assistant_sentence_entry.created Reply assistant sentence entry created
      reply_assistant_sentence_entry.destroyed Reply assistant sentence entry destroyed
      reply_assistant_sentence_entry.updated Reply assistant sentence entry updated
      reply_assistant_sentence_version.created Reply assistant sentence version created
      reply_assistant_sentence_version.destroyed Reply assistant sentence version destroyed
      reply_assistant_sentence_version.updated Reply assistant sentence version updated
      reply_assistant_version_permalink.created Reply assistant permalink version created
      reply_assistant_version_permalink.destroyed Reply assistant permalink version destroyed
      reply_assistant_version_permalink.updated Reply assistant permalink version updated
      role.created Role created
      role.destroyed Role destroyed
      role.updated Role updated
      security.updated Security settings updated
     session.created A user signed in. 
session.destroyed A user logout.
settings.updated Settings updated
survey.created Survey created
survey.destroyed Survey destroyed
survey.updated Survey updated
tag.created Tag created
tag.destroyed Tag destroyed
tag.updated Tag updated
team.created Team created
team.destroyed Team destroyed
team.updated Team updated
user.created Agent created
user.destroyed Agent destroyed
user.disconnected Agent disconnected
user.impersonated Agent impersonated
user.invited Agent invited
user.notifications_updated Agent's notifications updated
user.updated Agent updated}

events = []

doc.split("\n").each_with_index do |line,i|
    puts "LINE[#{i}] " + line
    parts = line.split(' ')
    if parts.length>2
        event = parts.shift
        desc = parts.join(' ')
        info = {
            event: event,
            desc: desc }
        puts MultiJson.dump(info)
        events.push event
    else
        puts "ERROR"
        break
    end
end 

puts MultiJson.dump(events)